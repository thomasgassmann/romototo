import requests
import logging
from selenium import webdriver
from bs4 import BeautifulSoup
from romototo.housing import Housing

# adapted from https://github.com/xiyichen/swissroll/blob/main/backend/zurich_housing/subscription/crawler.py

class Provider:
  
    def __init__(self, get):
        self._get = get

    def get_housing(self):
        return self._get()

def get_requests_headers():
    return {'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_3) '
                          'AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36'}

def get_selenium_headers():
    chrome_options = webdriver.ChromeOptions()
    chrome_options.add_argument('--headless')
    chrome_options.add_argument('--no-sandbox')
    chrome_options.add_argument('--disable-dev-shm-usage')
    chrome_options.add_argument('--headless')
    driver = webdriver.Chrome(executable_path='/usr/bin/chromedriver', options=chrome_options)
    return driver

def get_woko_type(address, rent, title):
    if 'studio' in title:
        return 'studio'
    if address.upper() == 'GSTEIGSTRASSE 18, 8049 ZÜRICH':
        return 'studio'
    if address.upper() == 'HIRZENBACHSTRASSE 4, 8051 ZÜRICH' and rent >= 800:
        return 'studio'
    if address.upper() == 'UETLIBERGSTRASSE 111/111A/111B, 8045 ZÜRICH' and rent >= 870:
        return 'studio'
    if address.upper() == 'AM WASSER 6/15, 8600 DÜBENDORF' and rent >= 1000:
        return 'studio'
    if address.upper() == 'ALTE LANDSTRASSE 98, 8702 ZOLLIKON' and rent >= 1000:
        return 'studio'
    if address.upper() == 'ALTSTETTERSTRASSE 183, 8048 ZÜRICH':
        return 'studio'
    return 'shared apartment'


def check_woko():
    logging.info('Checking woko')
    r = requests.get('http://www.woko.ch/en/nachmieter-gesucht', headers=get_requests_headers())
    soup = BeautifulSoup(r.text, 'lxml')
    items = soup.find("div", {"id": "GruppeID_98"}).find_all("div", {"class": "inserat"})
    logging.info(f'Woko has {len(items)} items')
    housings = []
    for item in items:
        link = 'http://www.woko.ch' + item.find('a', href=True)['href']
        listing_id = link.split('/')[-1]
        housing = Housing.build('woko', listing_id)
        title = item.find('h3').text
        logging.info('Checking woko listing: ' + title)
        rent = int(item.find("div", {"class": 'miete'}).find("div", {"class": "preis"}).text.split('.--')[0])
        start_date = item.find('table').find_all('tr')[0].find_all('td')[-1].text.strip().split(' from ')[-1]
        address = item.find('table').find_all('tr')[1].find_all('td')[-1].text
        listing_type = get_woko_type(address, rent, title.lower())
        r_d = requests.get(link, headers=get_requests_headers())
        s_d = BeautifulSoup(r_d.text, 'lxml')
        contact_person = s_d.find_all('table')[1].find_all('td')[1].text
        contact_email = s_d.find_all('table')[1].find_all('td')[3].text

        housing.set_general(title, link)
        housing.set_type(listing_type)
        housing.set_content(rent, start_date)
        housing.set_location(address)
        housing.set_contact(contact_person, contact_email)

        housings.append(housing)
    return housings

def check_living_science():
    logging.info('Checking living science')
    r = requests.get('http://reservation.livingscience.ch/en/living', headers=get_requests_headers())
    soup = BeautifulSoup(r.text, 'lxml')
    items = soup.find("div", {"class": "list scroll"}).find_all("div", {"class": "row status1"})
    housings = []
    for item in items:
        link = "http://reservation.livingscience.ch/en/living"
        # floor = item.find("span", {"class": "spalte1"}).text.split(':')[-1]
        rooms = item.find("span", {"class": "spalte4"}).text.split(':')[-1]
        rooms = float(rooms) if '.' in rooms else int(rooms)
        gross_rent = float(item.find("span", {"class": "spalte5"}).text.split(':')[-1].split(' ')[-1])
        start_date = item.find("span", {"class": "spalte6"}).text.split(':')[-1]
        apprnr = item.find("span", {"class": "spalte7"}).text.split(':')[-1]
        size = item.find("span", {"class": "spalte8"}).text.split(':')[-1].split(' ')[1]
        # housenr = item.find("span", {"class": "spalte10"}).text.split(':')[-1]
        charges = float(item.find("span", {"class": "spalte11"}).text.split(':')[-1].split(' ')[-1])
        housing = Housing.build('living_science', apprnr)
        housing.set_general('Living Science #' + apprnr, link)
        housing.set_number_of_rooms(rooms)
        housing.set_type('studio')
        housing.set_content(gross_rent, start_date)
        housing.set_size(size)
        housing.set_charges(charges)
        housings.append(housing)
    return housings

def check_wohnenuzhethz():
    logging.info('Checking ETH housing office')
    driver = get_selenium_headers()
    driver.get("https://wohnen.ethz.ch/index.php?act=searchoffer")
    language_button = driver.find_element('xpath', './/a[@class="language"]')
    language_button.click()
    not_furnished_checkbox = driver.find_element('xpath', './/input[@name="MoebilCheckbox_1"]')
    if not_furnished_checkbox.is_selected():
        not_furnished_checkbox.click()
    city_inputbox = driver.find_element('xpath', './/input[@name="Ort"]')
    city_inputbox.send_keys("Zürich")
    search_button = driver.find_element('xpath', './/input[@value="Search"]')
    search_button.click()
    driver.find_element('xpath', ".//select[@name='sortHow']/option[text()='Descending']").click()
    driver.find_element('xpath', ".//select[@name='MultiSortField1']/option[text()='No']").click()
    unfiltered_items = driver.find_elements('xpath', ".//table[@class='listing']/tbody/tr")
    housings = []
    for unfiltered_item in unfiltered_items:
        tds = unfiltered_item.find_elements('xpath', ".//td")
        if len(tds) < 12:
            continue
        apprnr = tds[1].find_element('xpath', ".//a").text
        logging.info('Checking ETH housing office listing: ' + apprnr)
        district = tds[4].text
        rent = "".join(tds[5].text.split("'"))
        listing_type = tds[6].text.lower()
        rooms = tds[7].text
        rooms = float(rooms) if '.' in rooms else int(rooms)
        size = tds[8].text
        start_date = tds[9].text
        end_date = tds[10].text
        furnished = tds[11].text
        link = 'https://wohnen.ethz.ch/index.php?act=detoffer&pid=' + apprnr
        housing = Housing.build('wohnenuzhethz', apprnr)
        if end_date:
            housing.set_end_date(end_date)
        housing.set_general('ETH / Uni Housing #' + apprnr, link)
        housing.set_location(district)
        housing.set_content(int(rent), start_date)
        housing.set_size(size)
        housing.set_type(listing_type)
        housing.set_number_of_rooms(rooms)
        housing.set_furnished(furnished)

        housings.append(housing)
    return housings

PROVIDERS = [Provider(check_woko)]

import pickle
import os
import logging
from typing import List
from romototo.crawl import PROVIDERS
from romototo.housing import Housing
from romototo.send import send_message
from time import sleep

def load_state(pickle_file):
    with open(pickle_file, 'rb') as f:
        return pickle.load(f)

def save_state(pickle_file, state):
    with open(pickle_file, 'wb') as f:
        pickle.dump(state, f)

def run_with_config(config):
    providers = PROVIDERS
    pickle = config['state']

    current = load_state(pickle) if os.path.exists(pickle) else []
    while True:
        new_items = []
        for provider in providers:
            try:
                housing: List[Housing] = provider.get_housing()
                logging.info(f'Current provider yielded {len(housing)} items')
            except Exception as e:
                logging.error(f'Error while crawling: {e}')
                continue

            for h in housing:
                if h.uuid() in current or h.rent() > config['max_rent']:
                    continue
                current.append(h.uuid())
                new_items.append(h)

        logging.info(f'Found {len(new_items)} new items')
        if len(new_items) > 0:
            send_message(config, 'romototo housing', build_body(new_items))
        logging.info(f'Sleeping for {config["crawl_interval"]} seconds')
        save_state(pickle, current)
        sleep(config['crawl_interval'])


def build_body(housing: List[Housing]):
    content = '<html><body>'

    content += '<p>Hi there, we have found some new housing opporunities:</p>'

    content += '<ul>'
    for h in housing:
        content += '<li>'
        content += h.to_html()
        content += '</li>'
    content += '</ul>'

    content += '</body></html>'
    return content

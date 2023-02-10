class Housing:

    _number_of_rooms = None
    _room_type = None
    _type = None
    _id = None
    _person = None
    _email = None
    _rent = None
    _start_date = None
    _end_date = None
    _charges = None
    _location = None
    _size = None
    _title = None
    _link = None
    _furnished = None

    def uuid(self):
        return f'{self._type}:{self._id}'

    def to_html(self):
        content = f'<a href="{self._link}">'
        content += f'{self._title} ({self._room_type}) CHF {self._rent}'
        if self._charges:
            content += f' (+ CHF {self._charges})'

        return content + '</a>'

    @staticmethod
    def build(type: str, id: str):
        housing = Housing()
        housing._type = type
        housing._id = id
        return housing

    def rent(self):
        return self._rent

    def set_contact(self, person: str, email: str):
        self._person = person
        self._email = email

    def set_content(self, rent: int, start_date: str):
        self._rent = rent
        self._start_date = start_date

    def set_charges(self, charges: int):
        self._charges = charges

    def set_location(self, location: str):
        self._location = location

    def set_size(self, size):
        self._size = size

    def set_general(self, title: str, link: str):
        self._title = title
        self._link = link

    def set_end_date(self, end_date: str):
        self._end_date = end_date

    def set_number_of_rooms(self, number: int):
        self._number_of_rooms = number

    def set_furnished(self, furnished: bool):
        self._furnished = furnished

    def set_type(self, type: str):
        self._room_type = type

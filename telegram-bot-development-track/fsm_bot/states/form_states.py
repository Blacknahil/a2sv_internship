from aiogram.fsm.state import State, StatesGroup

class About(StatesGroup):
    first_name = State()
    last_name = State()
    age = State()
    hobbies = State()

class Address(StatesGroup):
    city = State()
    street = State()
    house = State()

class Contact(StatesGroup):
    phone = State()
    email = State()

class Form(StatesGroup):
    about = About
    address = Address
    contact = Contact
    confirm = State()
    confirm_reject = State()

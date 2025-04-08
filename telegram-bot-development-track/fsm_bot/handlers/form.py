# djhsdhjhj
from typing import Dict
from aiogram import F, Router, html
from aiogram.filters import CommandStart,Command
from aiogram.types import Message, ReplyKeyboardMarkup, KeyboardButton, ReplyKeyboardRemove
from aiogram.fsm.context import FSMContext
from states.form_states import Form
import logging 
from aiogram.exceptions import StopPropagation


form_router = Router()  # This creates your router

@form_router.message(CommandStart())
async def command_start(message:Message,state:FSMContext) -> None:
    await state.set_state(Form.about.first_name)
    await message.answer(
        "Hello, What's your name?",
        reply_markup= ReplyKeyboardMarkup(
            keyboard=[
                [KeyboardButton(text='Skip')],
                [KeyboardButton(text="Cancel")],
            ],
            resize_keyboard=True,
        )
    )

@form_router.message(Form.about.first_name)
async def process_first_name(message:Message,state:FSMContext) ->None:
    await state.update_data(first_name=message.text)
    await state.set_state(Form.about.last_name)
    await message.answer(
        "What's your last name?",
        reply_markup = ReplyKeyboardMarkup(
            keyboard=[
                [
                    KeyboardButton(text="Skip"),
                    KeyboardButton(text="Go_back")
                    ],
                [
                    KeyboardButton(text="Cancel")
                ]
            ],
            resize_keyboard=True,
        )
    )

@form_router.message(Form.about.last_name)
async def process_last_name(message:Message,state:FSMContext) ->None:
    await state.update_data(last_name=message.text)
    await state.set_state(Form.about.age)
    await message.answer(
        "How old are you ?",
        reply_markup=ReplyKeyboardMarkup(
            keyboard=[
                [KeyboardButton(text="Skip"),KeyboardButton(text="Go Back")],
                [KeyboardButton(text="Cancel")],
            ]
        )
    )

@form_router.message(Form.about.age)
async def process_age(message:Message,state:FSMContext) ->None:
    await state.update_data(age=message.text)
    
    # to the next about.hobbies stage 
    await state.set_state(Form.about.hobbies)
    await message.answer(
        "What is your hobby?",
        reply_markup=ReplyKeyboardMarkup(
            keyboard=[
                [KeyboardButton(text="Skip"),KeyboardButton(text="Go Back")],
                [KeyboardButton(text="Cancel")],
            ]
        )
    )
    

@form_router.message(Form.about.hobbies)
async def process_hobbies(message:Message,state:FSMContext)->None:
    await state.update_data(hobbies=message.text)
    
    # lets go to the next stage of address(city,street,house)
    await state.set_state(Form.address.city)
    await message.answer(
        "What is the name of the city you reside ?",
        reply_markup=ReplyKeyboardMarkup(
            keyboard=[
                [KeyboardButton(text="Skip"),KeyboardButton(text="Go Back")],
                [KeyboardButton(text="Cancel")],
            ]
        )
    )
    
@form_router.message(Form.address.city)
async def process_city(message:Message,state:FSMContext) -> None:
    await state.update_data(city=message.text)
    
    # lets go the the next: address - street 
    await state.set_state(Form.address.street)
    await message.answer(
        "What street do you live in ?",
        reply_markup=ReplyKeyboardMarkup(
            keyboard=[
                [KeyboardButton(text="Skip"),KeyboardButton(text="Go Back")],
                [KeyboardButton(text="Cancel")],
            ]
        )
    )

@form_router.message(Form.address.street)
async def process_street(message:Message,state:FSMContext)->None:
    await state.update_data(street= message.text)
    
    # lets go to the next state: address(House)
    await state.set_state(Form.address.house)
    await message.answer(
        "What is your house number ?",
        reply_markup=ReplyKeyboardMarkup(
            keyboard=[
                [KeyboardButton(text="Skip"),KeyboardButton(text="Go Back")],
                [KeyboardButton(text="Cancel")],
            ]
        ) 
    )

@form_router.message(Form.address.house)
async def process_house(message:Message,state:FSMContext)-> None:
    await state.update_data(house=message.text)
    
    # lets go to the next stage: contact phone and email
    await state.set_state(Form.contact.phone)
    await message.answer(
        "What is your Phone number ?",
        reply_markup=ReplyKeyboardMarkup(
            keyboard=[
                [KeyboardButton(text="Skip"),KeyboardButton(text="Go Back")],
                [KeyboardButton(text="Cancel")],
            ]
        ) 
    )
@form_router.message(Form.contact.phone)
async def process_phone(message:Message,state:FSMContext)->None:
    await state.update_data(phone=message.text)
    
    #  lets go to the next contact info which is email
    await state.set_state(Form.contact.email)
    await message.answer(
        "What is your Email ?",
        reply_markup=ReplyKeyboardMarkup(
            keyboard=[
                [KeyboardButton(text="Skip"),KeyboardButton(text="Go Back")],
                [KeyboardButton(text="Cancel")],
            ]
        )  
    )

@form_router.message(Form.contact.email)
async def process_email(message:Message,state:FSMContext)-> None:
    await state.update_data(email=message.text)
    await state.set_state(Form.confirm)
    data : Dict[str,any]= await state.get_data()
    await message.answer(
        f"Please confirm your data:\n"
        f"Name: {html.quote(data['first_name'])}"
        f"{html.quote(data['last_name'])}\n"
        f"Age: {html.quote(data['age'])}\n"
        f"City: {html.quote(data['city'])}\n"
        f"Street: {html.quote(data['street'])}\n"
        f"House : {html.quote(data['house'])}\n"
        f"Phone: {html.quote(data['phone'])}\n"
        f"Email: {html.quote(data['email'])}\n",
        reply_markup=ReplyKeyboardMarkup(
            keyboard=[
                [KeyboardButton(text="Approve"),KeyboardButton(text="Disapprove")],
                [KeyboardButton(text="Go back")],
                [KeyboardButton(text="cancel")],
            ],
            resize_keyboard=True,
        )
    )
    

@form_router.message(Command("cancel"))
@form_router.message(lambda message: message.text and message.text.casefold() == "cancel")
async def cancel_handler(message: Message, state: FSMContext) -> None:
    """
    Allow user to cancel any action
    """
    current_state = await state.get_state()
    if current_state is None:
        return
    logging.info("Cancelling state %r", current_state)
    await state.clear()
    await message.answer(
        "Cancelled.",
        reply_markup=ReplyKeyboardRemove(),
    )
    raise StopPropagation


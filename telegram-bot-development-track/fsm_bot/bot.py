import os
from dotenv import load_dotenv
import asyncio
import logging
from handlers import form 
from aiogram.fsm.storage.memory import MemoryStorage
from aiogram.types import Message
from aiogram import Bot,Dispatcher


load_dotenv("../.env.local")
TOKEN = os.getenv("TELEGRAM_TOKEN")

# Configure logging
logging.basicConfig(level=logging.INFO)

async def main():
    if not TOKEN:
        raise Exception( "Token not found")
        
    bot = Bot(token=TOKEN)
    dp = Dispatcher(storage=MemoryStorage())
    dp.include_router(form.form_router)
    try:
        logging.info("Starting bot polling...")
        await dp.start_polling(bot)
    except Exception as e:
        logging.error(f"An error occurred while polling: {e}")
    finally:
        await bot.session.close()
        logging.info("Bot stopped.")


if __name__=="__main__":
    try:
        asyncio.run(main())
    except Exception as e:
        print(f"An error occurred: {e}")
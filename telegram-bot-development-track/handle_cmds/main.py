# djhjhksdhj 
import logging
import os
from dotenv import load_dotenv
# from telegram.ext import CommandHandler 
from telegram.ext import CommandHandler,MessageHandler,filters,Application
import re

# loading the telegram bot token 
load_dotenv(dotenv_path="../.env.local")
BOT_TOKEN = os.getenv("TELEGRAM_TOKEN")

if not BOT_TOKEN:
    raise ValueError("BOT_TOKEN is not set. Check your .env file.")


# set up the logging 

logging.basicConfig(format='%(asctime)s - %(name)s - %(levelname)s - %(message)s', level=logging.INFO)
logging.info("starting Nahom's Bot....")
async def start_command(update,context):
    await update.message.reply_text("Hello there! I\'m Nahom\'s bot. Good to see you here!")
    
async def help_command(update,context):
    await update.message.reply_text("Try typing what you have in mind and will do my best to help!")
async def custom_command(update,context):
    await update.message.reply_text("This is a custom command.")
async def handle_message(update,context):
    text = update.message.text.lower()
    logging.info(f"User ({update.message.chat_id}) says: {text}")

    # Regex-based responses
    if re.search(r"\bhello\b", text):
        await update.message.reply_text("Hi there! How can I assist you?")
    elif re.search(r"\bbye\b", text):
        await update.message.reply_text("Goodbye! Have a great day!")
    elif re.search(r"\bthanks\b|\bthank you\b", text):
        await update.message.reply_text("You're welcome! 😊")
    else:
        await update.message.reply_text(text)  # Default echo

# Handle unknown commands
async def unknown_command(update,context):
    await update.message.reply_text("Sorry, I don't recognize that command.")

async def echo_command(update,context):
    """This function echoes back what the user sends."""
    if context.args:
        await update.message.reply_text(" ".join(context.args))
    else:
        await update.message.reply_text("Usage: /echo <your message>")

async def error(update,context):
    # just logging the messaages 
    logging.error(f'Update {update} caused error {context.error}')


if __name__ == "__main__":
    # initalize the bot application 
    app = Application.builder().token(BOT_TOKEN).build()
    
    # commands 
    app.add_handler(CommandHandler("start",start_command))
    app.add_handler(CommandHandler("help",help_command))
    app.add_handler(CommandHandler("custom",custom_command))
    app.add_handler(CommandHandler("echo", echo_command))
    app.add_handler(MessageHandler(filters.COMMAND ,unknown_command))
    app.add_handler(MessageHandler(filters.TEXT & ~filters.COMMAND, handle_message))    # for any messages echo just the message  # Add a message handler for text messages
    app.add_error_handler(error)
    logging.info("Bot is polling ...")
    app.run_polling()

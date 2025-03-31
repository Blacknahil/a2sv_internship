from dotenv import load_dotenv
import os
import logging
from telegram import Update
from telegram.ext import Application, CommandHandler, ContextTypes

# Load environment variables
load_dotenv(dotenv_path=".env.local")
BOT_TOKEN = os.getenv("TELEGRAM_TOKEN")

# Initialize logging
logging.basicConfig(
    format="%(asctime)s - %(levelname)s - %(message)s", 
    level=logging.INFO
)

# Initialize the bot application
app = Application.builder().token(BOT_TOKEN).build()

# Define the start command handler
async def start(update: Update, context: ContextTypes.DEFAULT_TYPE):
    await update.message.reply_text("Hello! I'm your python-telegram-bot 🤖")

# Add the command handler
app.add_handler(CommandHandler("start", start))

# Run the bot with Long Polling
if __name__ == "__main__":
    logging.info("Bot is starting with Long Polling...")
    app.run_polling(allowed_updates=Update.ALL_TYPES)

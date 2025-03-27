from dotenv import load_dotenv
import logging
from telegram import Update
from telegram.ext import Application, CommandHandler, ContextTypes

import os

load_dotenv(dotenv_path =".env.local")
BOT_TOKEN = os.getenv("TELEGRAM_TOKEN")


# Initialize the bot application
app = Application.builder().token(BOT_TOKEN).build()

# Define the start command handler
async def start(update: Update, context: ContextTypes.DEFAULT_TYPE):
    await update.message.reply_text("Hello! I'm your python-telegram-bot 🤖")

# Add the command handler
app.add_handler(CommandHandler("start", start))

# Run the bot
if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    app.run_polling()



# //jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
# //jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
# //jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
# //jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsddjhfdsjhjhdsjhjksdjk
# //jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
# //jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
# //jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
# //jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
# //jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
# //jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
# //jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
# //jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
# //jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
# //jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
# //jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
# //jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
# //jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
# //jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
# //jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
# //jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
# # djkfsadjkfsjdk

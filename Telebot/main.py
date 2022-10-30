import telebot
from telebot import types

bot = telebot.TeleBot('5758161228:AAHbQpgvOnYIJa2WTitiFkMcG-jEm75DIVc')

@bot.message_handler(commands=['start'])
def start(message):
    mess = f'Hi, <b>{message.from_user.first_name}</b>'
    bot.send_message(message.chat.id, mess, parse_mode="html")

@bot.message_handler(content_types=['text'])
def get_user_text(message):
    if message.text == "Hi":
       bot.send_message(message.chat.id, "Hi! =)!", parse_mode='html')
    elif message.text == "id":
        bot.send_message(message.chat.id, f"Твой ID: {message.from_user.id}", parse_mode='html')
    elif message.text == "Photo":
        photo = open('fish.png', 'rb')
        bot.send_photo(message.chat.id, photo)
    elif message.text == "Say hi":
        voice = open('hi.m4a', 'rb')
        bot.send_voice(message.chat.id, voice)
    else:
        bot.send_message(message.chat.id, "I don't understand you", parse_mode='html')
        
@bot.message_handler(content_types=['photo']) 
def get_user_photo(message):
    bot.send_message(message.chat.id, 'Wow! Nice picture!')    
    
@bot.message_handler(commands=['website'])   
def website(message):
    markup = types.InlineKeyboardMarkup()
    markup.add(types.InlineKeyboardButton("Our website", url="https://golangify.com/"))
    bot.send_message(message.chat.id, 'Go to the website', reply_markup=markup)
     
@bot.message_handler(commands=['help'])   
def website(message):
    markup = types.ReplyKeyboardMarkup(resize_keyboard=True, row_width=1) #the numberr of buttons
    website = types.KeyboardButton('Website')
    start = types.KeyboardButton('Start')
    markup.add(website, start)
    bot.send_message(message.chat.id, 'Go to the website', reply_markup=markup)
        
bot.polling(none_stop=True)

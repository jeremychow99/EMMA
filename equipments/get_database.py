from pymongo import MongoClient
from dotenv import load_dotenv
import os

def get_database():

    DB_URI = os.getenv('DB_URI')
    DATABASE = MongoClient(DB_URI)

    return DATABASE['Equipment']


if __name__ == "__main__":
    dbname = get_database()
import os
from dotenv import load_dotenv

load_dotenv()

class Settings:
    WEATHER_API_URL = os.getenv("WEATHER_API_URL", "https://api.open-meteo.com/v1/forecast")
    LAT = os.getenv("LAT")
    LON = os.getenv("LON")
    RABBITMQ_URL = os.getenv("RABBITMQ_URL", "amqp://guest:guest@rabbitmq:5672/")
    QUEUE_NAME = os.getenv("QUEUE_NAME", "weather.data")
    FETCH_INTERVAL = int(os.getenv("FETCH_INTERVAL", 3600))

settings = Settings()

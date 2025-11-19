import asyncio
import schedule
import time
from app.fetcher import fetch_weather
from app.publisher import RabbitMQPublisher
from app.config import settings

publisher = RabbitMQPublisher()

async def job():
    try:
        data = await fetch_weather()
        publisher.publish(data)
    except Exception as e:
        print("Error:", e)

def main():
    schedule.every(settings.FETCH_INTERVAL).seconds.do(
        lambda: asyncio.run(job())
    )

    print(f"Weather collector running every {settings.FETCH_INTERVAL}s")

    while True:
        schedule.run_pending()
        time.sleep(1)

if __name__ == "__main__":
    main()

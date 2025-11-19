import httpx
from app.config import settings

async def fetch_weather():
    params = {
        "latitude": settings.LAT,
        "longitude": settings.LON,
        "current_weather": True,
        "hourly": "temperature_2m,relativehumidity_2m,windspeed_10m,precipitation_probability"
    }

    async with httpx.AsyncClient(timeout=10) as client:
        response = await client.get(settings.WEATHER_API_URL, params=params)
        response.raise_for_status()
        data = response.json()

    # Normalize the payload
    current = data["current_weather"]

    normalized = {
        "timestamp": current["time"],
        "temperature": current["temperature"],
        "wind_speed": current["windspeed"],
        "humidity": data.get("hourly", {}).get("relativehumidity_2m", [None])[0],
        "rain_probability": data.get("hourly", {}).get("precipitation_probability", [None])[0],
        "lat": settings.LAT,
        "lon": settings.LON
    }

    return normalized

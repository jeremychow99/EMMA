FROM python:3-slim

WORKDIR /app

COPY requirements.txt .

RUN python -m pip install --no-cache-dir -r requirements.txt

COPY ./ .

EXPOSE 8080

CMD [ "python", "controller.py" ]



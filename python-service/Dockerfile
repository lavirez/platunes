FROM python:3.9.6-alpine
# FROM hub.hamdocker.ir/library/python:3.8

# set environment variables
ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

RUN mkdir /code
WORKDIR /code

# install dependencies
RUN pip install --upgrade pip
COPY requirements.txt/ /code/
RUN pip install -r requirements.txt

COPY . /code/

CMD ["sh", "/code/migrate_run.sh"]

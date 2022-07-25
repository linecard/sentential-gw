FROM python:3.9

ADD poetry.lock pyproject.toml /src/
ADD src/ /src/

WORKDIR /src

RUN pip install poetry && \
    poetry config virtualenvs.create false && \
    poetry install --no-interaction --no-ansi

ENV LAMBDA_ENDPOINT=${LAMBDA_ENDPOINT}

CMD ["python", "main.py"]

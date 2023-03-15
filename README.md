# Url shortener

A Url shortener service 
## ENV setup
rename env.example to .env

```
GO_PORT= 
GO_ENV=
PG_HOST=
PG_USER=
PG_DB=
PG_PORT=
URL_SHORTNER_API_KEY=
SHORT_URL_BASE_URI=
```
## API Reference

#### Get short url

```http
  GET /api/v1/url/short
```

| query       | Type     | Description                |
| :--------   | :------- | :------------------------- |
| `url`       | `string` | **Required**. long url     |

| Header       | Type     | Description                |
| :--------   | :------- | :-------------------------  |
| `x-api-key` | `string` | **Required**. Your API key  |

#### Create short url

```http
  POST /api/v1/url/short
```

| body        | Type     | Description                |
| :--------   | :------- | :------------------------- |
| `url`       | `string` | **Required**. long url     |

| Header      | Type     | Description                |
| :--------   | :------- | :------------------------- |
| `x-api-key` | `string` | **Required**. Your API key |

#### Redirect to long url

```http
  GET /:shortUrl
```

| params           | Type     | Description                |
| :--------        | :------- | :------------------------- |
| `shortUrl`       | `string` | **Required**. shortUrl     |

## Database setup
Please review the file 'db.sql' to see all SQL queries. We're looking to set up a database trigger that will convert base 10 numbers to base 62 alphanumeric characters for shortening url

## Tech Stack

go,gin,bun,PostgreSQL


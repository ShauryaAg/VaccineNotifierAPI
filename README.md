<p align="center">
    <img alt="VaccineNotifier" src="https://user-images.githubusercontent.com/31778302/118365915-36097280-b5bc-11eb-871a-4b9ab9d13f05.png" width="100" />
</p>
<h1 align="center"> Vaccine Notifier API </h1>

This API allows users to subscribe to notification emails, if any vaccine slot is empty within their pincode.

### Features

- Users can subscribe for emails about vaccine availability
- Email confirmation before subscribing to avoid spamming
- Unsubscribe button in emails if user wishes to do so.
- Cron Job runs every 1.5 hours to check for availability

### Setup locally

- Clone the repo using `git clone https://github.com/ShauryaAg/VaccineNotifierAPI.git`
- Move into the project folder `cd VaccineNotifierAPI/`
- Create a `.env` file in the project folder

```
CURRENT_HOST=<Your Deployment Link>
PORT=<PORT>
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=postgres
POSTGRES_HOST=postgres
SECRET=<Secret>
SENDGRID_API_KEY=<Your Key>
```

##### Using Docker <a href="https://www.docker.com/"> <img alt="VaccineNotifier" src="https://www.docker.com/sites/default/files/d8/styles/role_icon/public/2019-07/vertical-logo-monochromatic.png" width="50" /> </a>

- Run using `sudo docker-compose up`

### **OR**

##### Using Golang <a href="https://golang.org/"> <img alt="VaccineNotifier" src="https://golang.org/lib/godoc/images/go-logo-blue.svg" width="50" /> </a>

- Install the dependecies using `go mod download`
- Run using `go run server.go`

### Deployment

- > This can't be deployed on **Heroku** since **Cowin public API** doesn't allow requests from anywhere outside India (at least from my experience) and **Heroku** doesn't allow us to choose Indian region.
- > Deploy on AWS / Digital Ocean / GCP / Any other cloud service
- > It is currently deployed on AWS

## Endpoint Usage

#### Auth

- `/api/auth/register`

  - _Allowed Methods:_ `POST`
  - _Accepted Fields:_ `{name, email, password, age, pincode, preferredVaccine}`
  - _Returns:_ `User Details`
  - Sends a confirmation mail to the user

- `/api/auth/login`

  - _Allowed Methods:_ `POST`
  - _Accepted Fields:_ `{email, password}`
  - _Returns:_ `{id, email, token}`
    > Make sure email is confirmed before logging in

- `/api/auth/user`

  - _Allowed Methods:_ `GET` `PATCH`

    - `GET`

      - _Authorization:_ `Bearer <Token>`
      - _Returns:_ `User details after update`

    - `PATCH`

      - _Accepted Fields:_ `User details to be updated`
      - _Returns:_ `User details after update`

- `api/auth/unsub`

  - _Allowed Methods:_ `POST`
  - _Authorization:_ `Bearer <Token>`
  - Unsubscribe user from getting notification mails

- `api/auth/reset_password`

  - _Allowed Methods:_ `POST`
  - _Accepted Fields:_ `{email}`
  - Sends a password reset email to the user

#### Notification

- `api/notifyall`

  - _Allowed Methods:_ `GET`
  - Sends a notification email to all the registered users
    > This endpoint was created for testing purposes and doesn't work in production

#### Token

- `/t/<token>`

  - _Allowed Methods:_ `GET`
  - View to verify email confirmation token

- `/u/<token>`

  - _Allowed Methods:_ `GET`
  - View to unsubscribe user from notification emails

- `/f/<token>`

  - _Allowed Methods:_ `GET` `POST`
  - View to reset user password

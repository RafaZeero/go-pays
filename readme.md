# Go pays

"Go pays" is a rest API made with Go using GIN that allows users to create accounts and make transactions for your account.

## Technologies

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![MySQL](https://img.shields.io/badge/mysql-%2300f.svg?style=for-the-badge&logo=mysql&logoColor=white)
![PlanetScale](https://img.shields.io/badge/planetscale-%23000000.svg?style=for-the-badge&logo=planetscale&logoColor=white)

## API Endpoints

## GET

`GetAccounts` [/accounts](#get-/accounts) <br/>
`GetAccountsByID` [/accounts/:id](#get=/accounts/:id) <br/>

## POST

`CreateAccount` [/accounts](#post-/accounts) <br/>
`MakeTransaction` [/accounts/:id/transaction](#post-/accounts/:id/transaction)<br/>

## PATCH

`UpdateAccount` [/accounts](#patch-/accounts)<br/>

## DELETE

`DeleteAccount` [/accounts](#delete-/accounts)<br/>

---

### GET /accounts

Get all users accounts in the DB. <br/>
See implementation [here](https://github.com/RafaZeero/go-pays/blob/5d6465e942796d247474438a3d9c118a36d2af0c/handler/accounts.go#L51).

**Response**

```
// On success fetching accounts:
[
	{
		"id": 1,
		"name": "Rafa",
		"balance": 15000.99,
		"createdAt": "2023-09-01T14:20:19Z",
		"updatedAt": "2023-09-01T14:20:19Z"
	},
  {
		"id": 2,
		"name": "Rayssa",
		"balance": 500,
		"createdAt": "2023-09-01T14:20:23Z",
		"updatedAt": "2023-09-01T14:20:23Z"
	}
]

// On error:
{
  "error" : (error message)
}
```

### GET /accounts/:id

Get one account by id. <br/>
See implementation [here](https://github.com/RafaZeero/go-pays/blob/5d6465e942796d247474438a3d9c118a36d2af0c/handler/accounts.go#L85).

**Response**

```
// On success:
{
	"data": {
		"account": {
			"id": 9,
			"name": "Nami",
			"balance": 8000000,
			"createdAt": "2023-08-31T13:26:02Z",
			"updatedAt": "2023-09-01T13:39:57Z"
		}
	},
	"message": "Account found",
	"success": true
}

// On error:
// * User not found
{
  "error" : "User not found or does not exists"
}
// * Others
{
  "error" : (error message)
}
```

### POST /accounts

Create a new account. <br/>
See implementation [here](https://github.com/RafaZeero/go-pays/blob/5d6465e942796d247474438a3d9c118a36d2af0c/handler/accounts.go#L11).

**Parameters**

| Name      | Required | Type             | Description                          |
| --------- | -------- | ---------------- | ------------------------------------ |
| `name`    | required | string           | Name of the account owner            |
| `balance` | required | number (float64) | Initial balance value in the account |

**Response**

```
{
	"data": {
		"balance": 991299349956.78,
		"id": 48,
		"name": "Riquinho rico"
	},
	"message": "Account created",
	"success": true
}
```

### POST /accounts/:id/transaction

Make a transaction in an account. Can be a deposit or withdrawl. <br/>
See implementation [here](https://github.com/RafaZeero/go-pays/blob/5d6465e942796d247474438a3d9c118a36d2af0c/handler/accounts.go#L224).

**Parameters**

**Response**

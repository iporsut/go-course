# Mrello (Trello Clone)

## Install go-migrate

```
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

## Features

- Register new user
- Login user
- Logout user
- All user share same single board, User can load board with 3 columns
- Board fixed to have 3 columns (To do, Doing, Done)
- User can create new card in To do column
- User can move card to each other columns
- Card has only ID, title and description
- Card has information about when it was created, when it was updated and by who
- Card has editing history

## Mission 1 Register new user

### Step

1. Create User table migration

```
> migrate create -ext sql -dir db/migrations -seq create_users_table
/Users/weerasak/src/learning/go-course/workshop/mrello/db/migrations/000001_create_users_table.up.sql
/Users/weerasak/src/learning/go-course/workshop/mrello/db/migrations/000001_create_users_table.down.sql
```

```
make migrate-create NAME=create_users_table
```

2. Add 000001_create_users_table.up.sql script

```sql
CREATE TABLE IF NOT EXISTS users (
        user_id UUID PRIMARY KEY,
        email TEXT UNIQUE NOT NULL,
        password TEXT NOT NULL,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL
);
```

3. Add 000001_create_users_table.down.sql script

```sql
DROP TABLE IF EXISTS users;
```

3. Run migrate up

```
migrate -database "$POSTGRESQL_URL" -path db/migrations up
```

4. Create type User struct in file user.go

```go
type User struct {
	ID        uuid.UUID
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
```

5. Create UserRepository interface in user.go

```go
type UserRepository interface {
        CreateUser(ctx context.Context, email string, passwordHash string) (*User, error)
        FindUserByEmailAndPasswordHash(ctx context.Context, email string, passwordHash string) (*User, error)
}
```

6. Implements Repository in user_repository.go in package pgrepository

```go
package pgrepository

import (
	"context"
	"database/sql"
	"mrello"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

// FindUserByEmailAndPasswordHash returns a user with the given email and password hash.
func (r *userRepository) FindUserByEmailAndPasswordHash(ctx context.Context, email string, passwordHash string) (*mrello.User, error) {
	row := r.db.QueryRowContext(ctx, `SELECT user_id, email, created_at, updated_at FROM users WHERE email = $1 AND password_hash = $2`, email, passwordHash)
	var user mrello.User
	err := row.Scan(&user.ID, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, mrello.WrapErr(err, mrello.ErrCodeNotFound, "user not found")
		}
		return nil, mrello.WrapErr(err, mrello.ErrCodeOther, "error finding user")
	}

	return &user, nil
}

// CreateUser creates a user with the given email and password hash.
func (r *userRepository) CreateUser(ctx context.Context, email string, passwordHash string) (*mrello.User, error) {
	row := r.db.QueryRowContext(ctx, `INSERT INTO users (user_id, email, password) VALUES ($1, $2, $3) RETURNING user_id,  created_at, updated_at`, uuid.NewString(), email, passwordHash)
	var user mrello.User
	err := row.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, mrello.WrapErr(err, mrello.ErrCodeOther, "error creating user")
	}
	user.Email = email

	return &user, nil
}
```

7. Define custom error for Mrello API in error.go

```go
package mrello

import "fmt"

type ErrorCode int

const (
	ErrCodeOther ErrorCode = iota
	ErrCodeNotFound
)

func (c ErrorCode) String() string {
	switch c {
	case ErrCodeNotFound:
		return "not found"
	case ErrCodeOther:
		return "other"
	default:
		return "unknown"
	}
}

type Error struct {
	Code    ErrorCode
	Message string
	Err     error
}

func (e *Error) Error() string {
	var unwrapStr string
	if e.Err != nil {
		unwrapStr = e.Err.Error()
	}
	return fmt.Sprintf("error code %s - message %s: %s", e.Code.String(), e.Message, unwrapStr)
}

func WrapErr(err error, code ErrorCode, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
```

8. Create handler for manage user registration request in handler.go

```go
package mrello

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UserRegistrationHandler(userRepo UserRepository) func(*gin.Context) {
	type Request struct {
		Email           string `json:"email" binding:"required"`
		Password        string `json:"password" binding:"required,min=8,max=64"`
		PasswordConfirm string `json:"password_confirm" binding:"required,eqfield=Password"`
	}

	return func(c *gin.Context) {
		var req Request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.Error(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		h, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			c.Error(err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		passwordHash := string(h)

		user, err := userRepo.FindUserByEmailAndPasswordHash(c.Request.Context(), req.Email, passwordHash)
		if err != nil && !IsErrCode(err, ErrCodeNotFound) {
			c.Error(err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		if user != nil {
			c.Error(err)
			c.AbortWithStatus(http.StatusConflict)
			return
		}

		if _, err := userRepo.CreateUser(c.Request.Context(), req.Email, passwordHash); err != nil {
			c.Error(err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		c.Status(http.StatusCreated)
	}
}
```

9. Mapping router with handler in router.go

```go
func Router(e *gin.Engine, userRepo UserRepository) {
	e.POST("/users/registration", UserRegistrationHandler(userRepo))
}
```

10. Create gin.Engine and manage repository dependency in main.go

## Mission 2 Login / Logout user

## Mission 3 Create new card into To do column

## Mission 4 Move card to other column

## Mission 5 Save card changed history

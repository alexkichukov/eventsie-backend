```
       Microservices:
     /       |       \
[Events]   [Auth]   [Categories]
```

### Relationships

**1:N**
A user can create many events.
An event can be created by 1 user.

**N:M**
A user can favourite/attend many events.
An event can be favourited/attended by many users.

```json
// User
{
  "username": "johndoe",
  "password": "john$pass#",
  "email": "john.doe@gmail.com",
  "createdEvents": ["xJnd12", "of32lM", "n0jHbA"],
  "favouriteEvents": ["pau0iD", "Mk12pr", "aPj4hg"],
}

// Event
{
  "_id": "Mk12pr",
  "title": "An example event",
  "date": "Wed, 02 Nov 2022 12:10:45 GMT",
  "location": {
    "address": "Some Street",
    "city": "Plovdiv",
    "postcode": "4029"
  },
  "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
  "tags": ["Development", "Plovdiv"],
  "category": "xPj0mm",
  "price": {
    "type": "free"
  }
}

// Category
{
  "_id": "xPj0mm",
  "title": "Development",
  "description": "Events about software development",
  "events": ["la80on", "bflj12", "SjD113"]
}
```

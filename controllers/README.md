# API documentation

## API table

| No  | Endpoint      | Verb     | Body fields (JSON, if any)                | Meaning                      | Cookies required |
| --- | ------------- | -------- | ----------------------------------------- | ---------------------------- | ---------------- |
| 1   | /auth/signin  | `POST`   | `email`, `password`                       | Sign in                      | No               |
| 1   | /auth/signup  | `POST`   | `email`, `password`                       | Sign up                      | No               |
| 1   | /auth/signout | `POST`   |                                           | Sign out                     | Yes              |
| 1   | /user/:id     | `GET`    |                                           | Get user info by user ID     | Yes              |
| 1   | /user/:id     | `DELETE` |                                           | Delete user by user ID       | Yes              |
| 1   | /class        | `POST`   | `ClassName`, `Description`                | Create a class               | Yes              |
| 1   | /class/:id    | `GET`    |                                           | Get a class info by class ID | Yes              |
| 1   | /class/:id    | `DELETE` |                                           | Delete a class by class ID   | Yes              |
| 1   | /report       | `POST`   | `ReportObjectID`, `ReportType`, `Message` | Create a report              | Yes              |
| 1   | /report       | `GET`    |                                           | Get all reports              | Yes              |

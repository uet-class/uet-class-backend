# API documentation

## API table

| No  | Endpoint           | Verb     | Body fields (JSON, if any)                | Meaning                            | Cookies required |
| --- | ------------------ | -------- | ----------------------------------------- | ---------------------------------- | ---------------- |
| 1   | /auth/signin       | `POST`   | `Email`, `Password`                       | Sign in                            | No               |
| 2   | /auth/signup       | `POST`   | `Email`, `Password`                       | Sign up                            | No               |
| 3   | /auth/signout      | `POST`   |                                           | Sign out                           | Yes              |
| 4   | /user/:id          | `GET`    |                                           | Get user info by user ID           | Yes              |
| 5   | /user/:id          | `DELETE` |                                           | Delete user by user ID             | Yes              |
| 6   | /class             | `POST`   | `ClassName`, `Description`                | Create a class                     | Yes              |
| 7   | /class/:id/student | `POST`   | `Id`                                      | Add student to a class by class ID | Yes              |
| 8   | /class/            | `GET`    |                                           | Get all user's classes             | Yes              |
| 9   | /class/:id         | `GET`    |                                           | Get a class info by class ID       | Yes              |
| 10  | /class/:id         | `DELETE` |                                           | Delete a class by class ID         | Yes              |
| 11  | /report            | `POST`   | `ReportObjectID`, `ReportType`, `Message` | Create a report                    | Yes              |
| 12  | /report            | `GET`    |                                           | Get all reports                    | Yes              |

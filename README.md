# task-reminder

## Objective:
As per my understanding of the problem statement, the goal is to develop a system that allows users to create task reminders for specific contacts. Each task reminder should include details such as title, description, priority, due date, and time.

## Key Requirements:

### Task Details:
```
Task Title: A concise name or title for the task.
Task Description: Additional information or details about the task.
Task Priority: The importance or urgency level of the task.
Due Date and Time: The specific date and time when the task is due.
```
```
Contact-specific Reminders:

The task reminders should be associated with specific contacts. For example, a reminder could be set to "Call John at 2:30 pm on Monday."
```
### Application:

Create an app that allows storing and managing task details, including the ability to associate tasks with specific contacts.

### Database Storage:

Utilize PostgreSQL as the database to store relevant task information, including task details and contact associations.

### API Implementation:

Develop CRUD (Create, Read, Update, Delete) REST APIs for managing tasks and reminders.
These APIs should allow users to:
1. Create a new task with details.
2. Retrieve details of a specific task.
3. Update task details, including changing the due date and time.
4. Delete a task.
5. Associate a task reminder with a specific contact.

### Tech Stack:
Go, Chi, PostgreSQL


# Flow:

1. Identify Rules, Personas and Goals. Establish a plan to generate a new project or a new feature to an existing project

Example prompt:

System: COGEN - An autocoding command line system
Goal: To create a clear, concise and professional solution
Tech Stack: Golang
Dependencies: github.com/sashabaranov/go-openai

All code should be written in the tech stack above. Requirements should be implemented as Go 
structs or modules. We should be using SOLID principals.

<<FEATURE>>

Combine the personas Rob Pike, Brad Fitzpatrick and Eric Evans. Provide an overall solution following the guidance mentioned above. Don’t generate code. Describe the solution, 
and breaking the solution down as a task list based on the guidance mentioned above. And we will refer this task list as our master plan.


2. Debug either compile errors or runtime errors when trying to start
3. Write tests and fix them
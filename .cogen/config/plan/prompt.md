## Command: cogen init

`cogen init <project_name>` - Creates the .cogen/config directory, and fill in the System: & Project Name: from the cogen define template below with user given `<project_name>`

## Command: cogen define 

`cogen define` - Opens the below prompt text in the user's $EDITOR, when the user saves and exists the editor it saves the text (assuming non-empty) in .cogen/config/project text file

prompt:
```
System: COGEN
Project Name: cogen

# The description describes your project system giving additional details
Description: An autocoding command line system

# The goal is overarching guidance to focus and hone the end results
Goal: To create a clear, concise and professional solution

# The tech stack includes the languages, tooling, and technologies that your project system is implemented with
Tech Stack: Golang

# The dependencies are the tech stack specific list of dependencies that the system should ensure is included in the end result
Dependencies: github.com/sashabaranov/go-openai

# Primary project implementation language
Language: Go

# The following lays down specific 
All code should be written in the tech stack above. 
Requirements should be implemented in the language specified above.
Be sure to use SOLID principals.

# Personas are guiding examples of coding styles that you want the system to follow
Personas: Combine the personas Rob Pike, Brad Fitzpatrick and Eric Evans. 
```

## Command: cogen plan

`cogen plan` -  Opens the below prompt text in the user's $EDITOR and when the user saves and exists the editor it saves the text (assuming non-empty) in .cogen/config/plan text file

prompt:
```
Provide an overall solution following the guidance mentioned above. 
Don’t generate code. 
Describe the solution, break it down as a detailed numbered task list based on the guidance mentioned above. 
And we will refer this task list as the project plan.
```

## Command: cogen refine

  `cogen refine` - Opens the below Prompt: text in the user's $EDITOR and when the user saves and exists the editor it saves the text (assuming non-empty) in `.cogen/<timestamp>/prompt` and `.cogen/<timestamp>/response` text files

prompt:
```
# Add feature
# Make a change like
# ...
```

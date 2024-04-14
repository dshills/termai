# termai

A terminal AI assistant.

## Features
- Default support for multiple AIs and models
	- OpenAI - GPT
	- Gemini - gemini
	- Mistral - mistral
	- Anthropic - Claude
	- Ollama - text generator models
- Specialized options for code generation
- Generate code with or without detailed explanations
- Generate source code for specific languages
- Custom language specific prompts

## Initialization

```none
$ go install github.com/dshills/termai

$ termai -init

$HOME/.termai.json configuration file created.
1) Open the file
2) Add your API keys
3) Mark models you wish to use as Active
4) Mark one model as Default (Can be overridden)
5) Add any langugae specfic prompts to the "prompts" section

```

## Usage

```none
Usage: termai [options] [query]
  -color
        Highlighted output
  -defaults
        Prints the default model
  -explain
        Explain the solution returned
  -ft string
        Use prompt extensions for a specific file type
  -help
        Print usage information
  -init
        Generate a default configuration file
  -list
        List available models
  -model string
        Model to use
  -opt-prompt
        Using the selected model try and optimize the prompt
  -opt-prompt-send
        Optimize the prompt and then use it
  -prompt
        Output the prompt without calling the AI
```

## Example Usage

### Basic usage
```none
$ termai Write a poem about a cat named bitty

Bitty the cat, small and sweet
Purring softly at my feet
With eyes as bright as stars above
She fills my heart with endless love

Her whiskers twitch, her tail does swish
As she curls up in a cozy dish
Content and calm, she naps all day
In her own little world, she'll forever stay

Bitty, my companion, my furry friend
On whom I can always depend
In her gentle purrs, I find peace
Forever grateful for her sweet release
```
### Prompt

- prompt: returns the prompt it would use without calling the AI
- ft: Set the file type to go
```none
$ termai -ft go -prompt Write a rope algorithm

I want you to act as a world-class software developer with decades of experience in go.
I will ask you for output, and you have to give me unique, expertly written work.
Write comments for all functions and data structures.
Return only code with no explainations except code comments.
Write a rope algorithm
```

### Code

- ft: Set the file type to go
- explain: Adds prompt to give a detailed explanation of the generated code

```sh
$ termai -ft go -explain Write a hello world application
...
Output: Hello world program with explinations of how it works

```



### Prompt Optimization

- opt-prompt will query the AI to improve the prompt
```none
$ termai -model gpt-4 -ft go -opt-prompt -prompt Write a rope algorithm

You are an expert in prompt engineering.
Rewrite this AI prompt to get the best results for code generation.
The text appearing inside of quotes is the prompt to be optimized.
"Act as a highly experienced software developer specializing in go Explain it to a highly experienmced go developer. Your work should be expertly written with unique code comments for all functions and data structures. Your task is to create fully functional and bug free code. Provide only code with comments and no explanations. Write a rope algorithm"
```

- opt-prompt-send will optimize the prompt and then use it
```none
$ termai -ft go -opt-prompt-send Write poem generator

Optimized Prompt: "Write a poem generator in Go, showcasing expertly written code with unique comments for all functions and data structures. Aim for fully functional, bug-free code tailored for a highly experienced Go developer."

...
Ouput: generated code

```

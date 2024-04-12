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

## Instalzation

```sh
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

```sh
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
```sh
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
```sh
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

```
Sure, here is a hello world application written in Go:

```go
package main

import "fmt"

// main function is the entry point of the program
func main() {
    // Println is a function that prints the given string to standard output
    fmt.Println("Hello, World!")
}
```

In this program, we have a `main` function which is the entry point of the program. Inside the `main` function, we use the `fmt.Println` function to print the string "Hello, World!" to the standard output.

I have chosen to use the `fmt.Println` function to print the output because it is the standard way to print output in Go. It is easy to use and provides a flexible way to format and print strings.

This approach is simple and straightforward, making it easy to understand for beginners and experts alike. It follows the standard conventions of Go programming, making it a reliable and efficient way to print output in a Go program.

```

### Prompt Optimization

```sh
$ termai -model gpt-4 -ft go -opt-prompt -prompt Write a rope algorithm

You are an expert in prompt engineering.
Rewrite this AI prompt to get the best results for code generation.
The text appearing inside of quotes is the prompt to be optimized.
"Act as a highly experienced software developer specializing in go Explain it to a highly experienmced go developer. Your work should be expertly written with unique code comments for all functions and data structures. Your task is to create fully functional and bug free code. Provide only code with comments and no explanations. Write a rope algorithm"
```

```sh
$ termai -ft go -opt-prompt-send Write poem generator

Optimized Prompt: "Write a poem generator in Go, showcasing expertly written code with unique comments for all functions and data structures. Aim for fully functional, bug-free code tailored for a highly experienced Go developer."

package main

import (
        "fmt"
        "math/rand"
        "time"
)

// Word struct to hold word options for generating a poem
type Word struct {
        noun    []string
        verb    []string
        adjective []string
}

// GeneratePoem generates a random poem using the provided words
func GeneratePoem(words Word, numLines int) {
        for i := 0; i < numLines; i++ {
                line := ""
                for j := 0; j < 4; j++ {
                        // Randomly select a noun, verb, or adjective to construct a line
                        switch rand.Intn(3) {
                        case 0:
                                line += words.noun[rand.Intn(len(words.noun))] + " "
                        case 1:
                                line += words.verb[rand.Intn(len(words.verb))] + " "
                        case 2:
                                line += words.adjective[rand.Intn(len(words.adjective))] + " "
                        }
                }
                // Print the line of the poem
                fmt.Println(line)
        }
}

func main() {
        // Seed the random number generator
        rand.Seed(time.Now().UnixNano())

        // Define a set of words for generating the poem
        words := Word{
                noun:    []string{"moon", "sun", "star", "tree", "flower"},
                verb:    []string{"shines", "dances", "whispers", "sleeps", "blooms"},
                adjective: []string{"bright", "beautiful", "peaceful", "vibrant", "serene"},
        }

        // Generate a 4-line poem using the provided words
        GeneratePoem(words, 4)
}
```

# Interviewing for Frontend Engineers

### Before interview be clear these below

- What do you like about your current job?
- What sort of things are you looking for in your next role?
- Which companies have appealing engineering cultures? What makes it appealing?

## Application

### candidate

- One page
- Skills
- Experience
- Technology familiarity
- Education
- Accomplishments

### interviewer

- Which roles does this person fit?
- What sort of experience is required?
- Is there any bias in your selection process?

## Initial call

### candidate

- What do you do currently?
- What are some projects you’ve worked on recently?
- What are you looking for in your next role?
    - work with people that are smarter than you
    - solve chanllege projects
    - big impact in my company
- Why do you want to work for ______?
- What is your availability for the next steps?

### interviewer

- How many steps are in the interview process? How long does it generally take?
- How big is your engineering team?
- Which team would I be interviewing for?
- What is the culture like?
- Who are you competitors?
- What sort of projects would I work on?

### Prescreen

- What is the difference between const, let, and var?

<details>
    <summary>Details</summary>

const, You can't point a const to a different pointer. But you can modify that, you can add objects, or you can add properties to an object, things from array. We can't change that pointer.
let, you can change the pointer, but it's only gonna be scoped to whatever the closure is.
var, hoisted to the top. So const and let, if you try to access them before they're gonna throw a reference error, var will just throw undefined.
</details>

- Explain prototypical inheritance.

<details>
    <summary>Details</summary>
JavaScript has a prototype, a baseline prototype of the object. Everything inherits from the object, so you can do all sorts of things. Everything in JavaScript has a prototype, it has a baseline object that it inherits from. And when you create a new object based on the other object, you can either inherit all those properties, which you will be default, or you can overwrite them with your own and so on, and so forth.
</details>

- What is 'this' mean in JavaScript?

<details>
    <summary>Details</summary>
The global context of everything that is available to access. So all the objects and functions are available to you that are not locally defined.
</details>

- What is the data structure of the DOM?

<details>
    <summary>Details</summary>
Tree
</details>

- What is a Stack and a Queue? How would you create those data structures in JavaScript?

<details>
    <summary>Details</summary>
LIFO, FIFO, Last in, First Out is the stack, First in, First Out as a cue.
It is an array pop and push shift.
</details>

- How can you tell if an image element is loaded on a page?

<details>
    <summary>Details</summary>
There's an onload element of Images You can just say is it on, has it loaded?
</details>

- What is call() and appy()?

<details>
    <summary>Details</summary>
All I would wanna know on this question is their ways of changing the scope of the calling function. Yeah, and call is a series of arguments and applies an array of arguments. Nice, we don't necessarily need to apply as much anymore, because we have props or we have array spreading, things like that.
</details>

- What is event delegation and what are the performance tradeoffs?

<details>
    <summary>Details</summary>
So traditionally, if you have event handlers in HTML, you could apply an event handler to every single element you wanna have. Or using event delegation, you could say I wanna have one event listener, and that's at the top.
And when you click on something, it just bubbles up to the parent that handles the event. That's an event delegation. I would want someone to know this cuz event listeners are really expensive on a page cuz every time it renders, gotta be like did something happen? Did something happen?
So it's better to have one event handler versus 60 for performance reasons. Event delegation. If you said something about bubbling, that's also good to use.
</details>

- What is a Worker? When would you use one?

<details>
    <summary>Details</summary>
A worker is something you would use in a browser to offload computationally expensive work. Three different thread cuz JavaScript is single threaded, if you have something that's like tactically prime to 10,000, numbers, something you want to do that and workers you're not blocking the UI, cuz there's only one thread in JavaScript.
</details>

## Code test

### candidate

- Make your code as readable as possible
    - Comment your code
    - Don’t over complicate the architecture
- Don’t import too many libraries
- If you have time, add unit tests
- Ask questions!

### interviewer

- Make the problem as straightforward as possible(take 2 minutes to undersatnd)
- Be honest with the time constraints(take 2~3 hours)
- Have a code review checklist

<details>
    <summary>Example</summary>

|Average|Good|Exceptional|
|---|---|---|
|Application starts properly|Code is well documented|Modular architecture designed for extensibility|
|3/5 requirements complete|All requirements are complete|Created unit and integration tests|
||No errors are thrown in the console||

</details>

### Big O

- Big Omega() - _best case_
- Big Theta - _average case_
- Big O - _worst case_

count loop.

## Phone screen

### candidate

- Ask questions
- Talk out your solution
- Get comfortable with the environment

### interviewer

- Are you in a quiet area?
- Is the problem well worded?
- Does the candidate know what the requirements and restrictions are?
- Did you leave time for questions at the end?

- [example](https://codepen.io/jemyoung/pen/JjjyBRZ/0817cc37159377752b6cd9bf70d40883?editors=1011)
- [solution](https://codepen.io/jemyoung/pen/wvvpGRx)

## On-site

### candidate

- Practice writing code without a computer
- Go over general sample problems
- Ask your friends to test you
- Try to ask what the style of technical questions will be

### interviewer

- Make sure questions are relevant to the role
- Have backup interviewers
- Book the room for the day
- Allow time for breaks (bathroom, lunch, etc)

## Result

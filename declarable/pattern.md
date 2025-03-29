```
The := operator can do one trick that you cannot do with var: it allows you to assign
values to existing variables too. As long as at least one new variable is on the lefthand
side of the :=, any of the other variables can already exist:
x := 10
x, y := 30, "hello"
Using := has one limitation. If you are declaring a variable at the package level, you
must use var because := is not legal outside of functions.
How do you know which style to use? As always, choose what makes your intent
clearest. The most common declaration style within functions is :=. Outside of a
function, use declaration lists on the rare occasions when you are declaring multiple
package-level variables.
In some situations within functions, you should avoid :=:
• When initializing a variable to its zero value, use var x int. This makes it clear
that the zero value is intended.
• When assigning an untyped constant or a literal to a variable and the default type
for the constant or literal isn’t the type you want for the variable, use the long var
form with the type specified. While it is legal to use a type conversion to specify
the type of the value and use := to write x := byte(20), it is idiomatic to write
var x byte = 20.
• Because := allows you to assign to both new and existing variables, it sometimes
creates new variables when you think you are reusing existing ones (see “Shad‐
owing Variables” on page 68 for details). In those situations, explicitly declare all
your new variables with var to make it clear which variables are new, and then
use the assignment operator (=) to assign values to both new and old variables.
While var and := allow you to declare multiple variables on the same line, use this
style only when assigning multiple values returned from a function or the comma ok
idiom
You should rarely declare variables outside of functions, in what’s called the package
block (see “Blocks” on page 67). Package-level variables whose values change are a
bad idea. When you have a variable outside of a function, it can be difficult to track
the changes made to it, which makes it hard to understand how data is flowing
through your program. This can lead to subtle bugs. As a general rule, you should
only declare variables in the package block that are effectively immutable.
You might be wondering: does Go provide a way to ensure that a value is immutable?
It does, but it is a bit different from what you may have seen in other programming
languages. It’s time to learn about const.
```


# API CHANGES

- [x] make a funciton to find boilerplate files.
- [x] make a function to list boilerplate files within the tree.
- [x] make it call the boilerplate page
- [] make --new work
- [x] use custom Delims <bplate>  </bplate>. {b  }


# MAIN

- [] move everyting into the main package. idk for the heck of it.
- [] write a test for the page builder.
- [] the main command accepts a file path to the template.
- [] make the sub tests their own functions

- [] write some basic readme and review the internal command documentation.
- [] build and executable and test it out locally.
- [] figure out how to publish it and install it via go package manage.
- [] write a readme.
- [] release the application

- [] std in and std out.

# OTHER

- [x] figure out a way to pass userdefined values

- [] figure out if go templating language is suitable.
    - I think I wast it to leave in place any tags that don't have relavant data.
    - Actaully idk. It might be fine, for my usecases.

# FEATURES

- [] add file counter slug feature with 001.md 002.md
- [] add default command args to header
- [] flag overide options.
- [] custom path
- [] custom slug pattern
- [] user defined command options.
- [] counter
    - the incremnt the filename by one.
    - works with slug { count }
    - takes a decimal place for 002
    - loop over fileExists() with each value until we can create the file.
- [] date string.
    - accepts a formate string %Y-%m-%d
- [] add some colors to the command.
- [] refine the log messages. maybe just use fmt.print instead.



# ALPHA 

- [x] create a command for each file in the _boilerplate folder
- [x] the user runs the new page command.
- [x] Open the boiler plate file associated with command
- [x] Extract out the header options from the boilerplate file
- [x] Extract out the template from the boilerplate file without the header.
- [x] Write the first test
- [x] parse the header options
- [x] Mix the default options with then header options. (do flag options later)
- [x] check for missing options. like path.
- [x] Template the boilerplate with the options
- [x] write the new file
DONE


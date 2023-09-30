

# Searchable

any file with a name beginning with .boilerplate
and any file within the .boilerplate/ folder.

template file of
post.md and notes and wiki/facts

will search
./post.md
.boilerplates/post.md
post.md/.boilerplate
post.md/.boilerplate.*

./notes (if its a file then use the file)
notes/.boilerplate (if its a folder check for .boilerplate files)
notes/.boilerplate.*
.boilerplates/notes (if its neither a file or a folder, check if a file exists in the .boilerplates folder)

wiki/facts (if its a file then use it as the template)
wiki/facts/.boierplate (if its a folder, then check for .boilerplate file)
wiki/facts/.boilerplate.*
.boierplates/wiki/facts (if neither then check if its in the .boilerplates folder)

# LIST Methods

recursivley search for any .boilerplate files and .boilerplate.* files.
and recursively list all files within the .boilderplates folder.

Don't list any files. that 

# new

## strat 1 be smart. pick the best place
path exists as a file > error, that file already exists.
path esists as a folder > notes/.boilerplate creaste a .boilerplate file
path doesn't exist, and its Dir does.
notes > notes/.boilerplate (if notes is a folder then create a notes/.boilerplate file)
notes > .boilerplates/notes (if it dosne't exist, then create a file in the .boilerplates folder)

## strat 2 be dumb. always pick the .boilerplates folder.
.boilerplates/path already esists > error, that boilerplate aready esists
.boilerplates/path is a folder > create the boilerplate file without complaint.
.boilerplates/path doesn't exist. > create it without complaint.

wiki/facts > wiki/facts/.boilerplate
wiki/facts > .boilerplates/wiki/facts

post.md > post.md/.boilerplate
post.md > .boilerplates/post.md

# path

.boilerplates/wiki/facts > wiki/fasts?
.boilerplates/wiki/facts > wiki/?

wiki/facts/.boilerplate > wiki/facts/

.boilerpates/post.md > post.md/
_posts/.boilerplate > _posts

# STDin

Treat std in as the template

# STDout


# API AS SUB EXAMPLES

bplate
list all the boilerplates in the tree

bplate list
explicitly list all the boilerplates in the tree

bplate new <path> [...args]
create a new boilperplate file

bplate create <path> [...args]
create a new file from a boilerplate


# API AS FLAG EXAMPLES

bplate
list all the boilerplates in the tree

bplate --list
explicitly list all the boilerplates in the tree

bplate --new <path> [...args]
create a new boilperplate file

bplate <path> [...args]
create a new file from a boilerplate



# API USEAGE EXAMPLES

bplate
no boilerplates found.
create a new one with: `bplate --new path/to/boilerplate.md`

bplate --new .boilerplates/post.md
Created new `.boilerplates/post.md` template file.

bplate --new notes/
Created new `notes/.boilerplate` template file.


bplate
$ bplate .boilerplates/post.md <title> [tags]
> _posts/{{ date }}-{{ title | slugify }}.md

$ bplate notes/.boilerplate <title>
> {{ path }}/{{ date }}-{{ title | slugify }}.md


bplate .boilerplates/post.md
Missing requiered <title> argment.


bplate .boilerplates/post.md "Foo Bar"
or
bplate post.md "Foo Bar"
Created `_post/yyyy-mm-dd-foo-bar.md` from `.boilerplates/post.md` boilerplate.


# API THOUGHTS

bplate --new path/to/new/file.md [...args]

bplate path/to/template.md --slug="{{ count }}.md" [...args]

bplate path/to/template.md [...args]


bplate --list

posts_boilerplate.md


bplate .boilerplates/post.md "Foo Bar"

Created `_` from `%v` boilerplate.", slug,name)

bplate --new path/to/new/file.md [...args]

bplate path/to/template.md --slug="{{ count }}.md" [...args]

bplate path/to/template.md [...args]


bplate --list

posts_boilerplate.md




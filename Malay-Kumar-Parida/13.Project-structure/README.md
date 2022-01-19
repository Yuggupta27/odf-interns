#assets#

We can put all other assets to go along with our project (images, logos, etc).


#cmd#

This folder contains the main application entry point files for the project, with the directory name matching the name for the binary.

It's common to have a small main function that imports and invokes the code from the /internal and /pkg directories and nothing else.

#docs#

All the design and user documents are stored here

#examples#

Examples for the applications and/or public libraries.


#internal#

This package holds the private library code used in the project, it is specific to the function of the project and not shared with other services.

#pkg#

Library code that's ok to use by external applications, Other projects will import these libraries expecting them to work

#vendor#

When using modules, the go command typically satisfies dependencies by downloading modules from their sources into the module cache and then loading packages from those downloaded copies. Vendoring may be used to ensure that all files used for a build are stored in a single file tree in the main module’s root directory containing copies of all packages needed to build and test packages in the main module

# Writing web application in GO
## https://golang.org/doc/articles/wiki/

Here are some simple tasks you might want to tackle on your own:

* Store templates in tmpl/ and page data in data/. :white_check_mark:
* Add a handler to make the web root redirect to /view/FrontPage. :white_check_mark:
* Spruce up the page templates by making them valid HTML and adding some CSS rules.
* Implement inter-page linking by converting instances of [PageName] to <a href="/view/PageName">PageName</a>. (hint: you could use regexp.ReplaceAllFunc to do this) :white_check_mark:
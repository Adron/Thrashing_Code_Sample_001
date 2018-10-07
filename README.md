# Thrashing Code Repo X

I put together this project in recent [video stream starting at 1:32:40](https://youtu.be/sg4Nnnb-Vvc?t=5560).

Here's the time slices with respective topics to check out:

<a  href="/watch?v=sg4Nnnb-Vvc&amp;t=10s">0:10</a> - Music Intro.

<a  href="/watch?v=sg4Nnnb-Vvc&amp;t=184s">3:04</a> - Introduction, recovering from it being morning on a Saturday, coffee, and getting some environment settings fixed.

<a  href="/watch?v=sg4Nnnb-Vvc&amp;t=345s">5:45</a> - Reviewing what has been done up to this point with Colligere.

<a  href="/watch?v=sg4Nnnb-Vvc&amp;t=460s">7:40</a> - Stating objectives for the day. Working on getting started with the schema file for the project.

<a  href="/watch?v=sg4Nnnb-Vvc&amp;t=675s">11:15</a> - Going over what we get when adding commands to a Cobra CLI. Where I re-stumble upon the fact I wrote out the commands as nested to the config sub-command. i.e. `colligere config set` and `colligere config create` etc. During this section I also go ahead and step through the Go lint output problem list and resolve the items.

<a  href="/watch?v=sg4Nnnb-Vvc&amp;t=1390s">23:10</a> - I run into a situation where Visual Studio Code has gotten out of sync with the Go Libs that are actually loaded in the Vendor directory, but I delete the directory, run `dep ensure` again and realizing it is just going to stay out of sync (anybody got a solution to this?) continue onward. But it's fascinating and helpful to know that this happens.

<a  href="/watch?v=sg4Nnnb-Vvc&amp;t=1700s">28:20</a> - I opt to start working on the `colligere config create` command to start fleshing out what needs to be set for JSON and written to file.

During this time I fight a lot with the permissions, how to set it, whether to use the const variables or pass the acual four digit permission code. In the end I finally end up just passing the 4 digits to create the directory and file. I also start using the Go user library and it's objects to pull user information for the project.
<a  href="/watch?v=sg4Nnnb-Vvc&amp;t=4855s">1:20:55</a> Thrashing Code Episode 3 &amp; bio break.

<a  href="/watch?v=sg4Nnnb-Vvc&amp;t=5560s">1:32:40</a> - Creating a new application with Jetbrains Goland to show off how to use the Go core libraries around the user, JSON, and some basic file creation and writing.

<a  href="/watch?v=sg4Nnnb-Vvc&amp;t=5800s">1:36:40</a> - At this point I start adding some basic code to pull the current user and collect some information about that user.

<a  href="/watch?v=sg4Nnnb-Vvc&amp;t=6042s">1:40:42</a> - I extract the error code using the Goland refactoring feature and setup a `func check()` for error checking. That cleans up the inline code a bit.

<a  href="/watch?v=sg4Nnnb-Vvc&amp;t=6190s">1:43:10</a> - Here I add to the user data retrieved some environment variables to that list of collected data. I also cover again, as I have a number of times, how the environment variables are pulled in IDE versus user session versus out of IDE.

<a  href="/watch?v=sg4Nnnb-Vvc&amp;t=6466s">1:47:46</a> - Now I add a file exists check and start working on that logic.

<a  href="/watch?v=sg4Nnnb-Vvc&amp;t=6536s">1:48:56</a> - Props to Edd Turtle on a solid site on Go. Here's the blog entry, and respectively the path to more of Edd's material @ https://golangcode.com/. Also, Edd seems like a good guy to follow [@eddturtule](https://twitter.com/eddturtle).

<a  href="/watch?v=sg4Nnnb-Vvc&amp;t=7293s">2:01:33</a> - Starting the JSON Work here to marshal and unmarshall.

<a  href="/watch?v=sg4Nnnb-Vvc&amp;t=8793s">2:26:33</a> - At this point I push the code up to Github using the built in Goland VCS features. I realize I've named the repo "adron" by accident so I rename it, close Goland, and then clone the code back down locally with Goland's VCS features. It's kind of interesting to see Goland go through the 2-factor auth for this too.

<a  href="/watch?v=sg4Nnnb-Vvc&amp;t=10736s">2:58:56</a> - The Seattle Thrashing Code outtro.</yt-formatted-string></div>

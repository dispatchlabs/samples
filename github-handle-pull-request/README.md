# ![](https://storage.googleapis.com/material-icons/external-assets/v4/icons/svg/ic_call_merge_black_24px.svg) Merge Code From Contributors (aka PR / Pull Request)

Example for `https://github.com/dispatchlabs/disgo`


- Get the **NEW** changes and **SEE** them from the contributing user
	```shell
	git clone https://github.com/dispatchlabs/disgo.git

	cd disgo

	git checkout -b pr-feature-1

	git pull https://github.com/__USER__/disgo.git master

	git checkout master
	git merge --no-ff --no-commit pr-feature-1
	```
- Verify then `git commit -m "..."` followed by `git push` 
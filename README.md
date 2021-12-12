# jjversion-gha-output

jjversion-gha-output is a dependency of [jjliggett/jjversion-action](https://github.com/jjliggett/jjversion-action), which is a custom composite GitHub action for versioning git projects.

jjversion-gha-output utilizes [jjliggett/jjversion](https://github.com/jjliggett/jjversion) to calculate versions for git repositories. The application uses the calculated version to output values for GitHub workflows, with the ```set-output``` workflow command, as detailed here: <https://docs.github.com/en/actions/learn-github-actions/workflow-commands-for-github-actions>

For information on how jjversion works, refer to the jjversion repository. For output of the jjversion-action GitHub action and other information on the composite action, refer to the jjversion-action repository.

## Licensing

Licensing for this application can be found at: [LICENSE.md](LICENSE.md).

The jjversion-gha-output license applies to all parts of the application that are not
externally maintained libraries and dependencies.

The main dependency of jjversion-gha-output is <https://github.com/jjliggett/jjversion>. In addition to jjversion, there are many indirect dependencies. Details on these dependencies can be found in the jjversion repository at <https://github.com/jjliggett/jjversion/blob/root/docs/NOTICE.md> along with the related files in the notices directory.

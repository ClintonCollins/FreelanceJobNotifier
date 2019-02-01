# Freelance Job Notifier

### What does it do?

Freelance Job Notifier checks certain enabled freelancer websites for new jobs matching your
defined queries. It can then alert you via enabled notifications.

### Installation

##### Install through build artifacts

You can download the latest build artifact and start running it right away. Binaries are pre-built
for Windows 64bit and Linux 64bit.

1. Download the latest artifacts.
2. Extract the artifact archive.
3. Copy your binary from the dist folder. Either `dist/linux64/freelance_job_notifer`
 or `dist/windows64/freelance_job_notifier`
4. Edit the config.toml file to your desired settings.
5. Run the binary. `./freelance_job_notifier` on linux or just double clicking it on Windows.

##### Build it yourself

You can clone this repository and just build it yourself. You'll need to download Go for your system.
https://golang.org/

Then switch to the source directory and run: `go build`

Rename default_config.toml to config.toml and edit it.

After that follow the steps above starting at #4

### Extra Information

Freelance Job Notifier is by no means perfect, but it's a good learning experience for me. 

### Future

Freelance Job Notifier needs documentation when I'm not as lazy. I plan on adding more notification avenues
such as Slack, SMTP, IRC, etc.

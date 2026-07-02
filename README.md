# 🔎 Verify 

Verify is a lightweight wrapper for the [Boot.dev CLI](https://github.com/bootdotdev/bootdev). It intercepts bootdev run commands to provide numbered test cases, improving the student experience.

---

## 🏋️ Motivation
Submitting your `run` cases with the `bootdev run <id>` command is easy, but then you have the arduous task of _verifying_ which of the run tests actually passed. In the standard bootdev cli output, test cases are not numbered and outputs can sometimes be very long, making your search **painful**. 

Verify fixes this, giving you numbered outputs on your test cases, making it easier to locate test headers and _verify_ result, saving you time _and_ frustration.

Happy with your test results? Great! Submitting the graded version is just a 'y' key away. 

--- 

## ✅ Requirements 
- [go](https://go.dev/doc/install) (developed and tested on v1.26.4)
- [bootdev cli tool](https://github.com/bootdotdev/bootdev) (developed and tested on v1.29.6)

---

## 🚀 Getting Started

### Install Verify
```bash
git clone https://github.com/breaking-boot/verify.git
cd verify
go install .
```

### Setup the Verify Alias
For Verify to run properly, you need to add a new alias to your shell configuration file (e.g., `.bashrc`, `.zshrc`, or `.profile`):
```bash
#verify alias for the bootdev cli tool
alias bootdev='verify' #comment out or delete this alias if you want to disable Verify, forcing tests to exclusively run through the official bootdev cli.
```
This tells the terminal to run `verify` whenever you run the `bootdev` cli tool.

Save the file and restart your terminal for the change to take effect.

---

## Usage Notes
Verify accepts any commands that the official bootdev cli accept, but our custom logic is only applied to commands in the format of:
```bash
bootdev run <id> # run
```

Entering other command arguments/flags may have unintended consequences. Typically any issues/errors are still caught by the original bootdev cli and output to the terminal. You may still see a cli output saying that Verify has run, even though no modification was actually applied to the original bootdev output.

---

## 🌱 Future Plans
- Numbered run-submit tests ❌ (abandoned, for now)
- Improved visual indicators
- Quick submit after run case ✅
- Check for expected results in run cases

---

## 📄 Contributing

If you have suggestions or fixes, feel free to open an issue or pull request 👍

---

## 📜 License

MIT License

---

## 👤 Author

[jeffschoe](https://github.com/jeffschoe)

---

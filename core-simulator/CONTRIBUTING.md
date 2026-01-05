# Contributing to open-exposure


Thanks for checking out the **open-exposure** project. We are building an open and modern 5G Network Exposure Function (NEF) to help telco engineers, verticals, and cloud-native developers integrate and innovate faster.

Whether you're fixing a bug, adding a feature, improving the docs, or just trying things out, you're very welcome here.

---

## Why contribute?

open-exposure is designed to:
- Expose 5G Core network capabilities (like monitoring events, QoS, traffic influence, and more) to applications in a safe and developer-friendly way.
- Provide tools and simulators to help integration with telco cloud providers.
- Create a community-driven alternative to expensive, vendor-locked NEF solutions.

Your contribution helps us make 5G and beyond innovation accessible to everyone.

---

## How to contribute

You can contribute in many ways:
- Fix typos or improve documentation
- Report bugs or suggest enhancements
- Add support for more 3GPP APIs or events
- Improve 3GPP procedures and workflows implementations
- Add new components to the existing implementation

Before working on a **big feature**, please open a GitLab issue to discuss your idea. Let's avoid duplicated work and align on the best approach.

---

## Reporting Issues

Found a bug? No problem!
1. Check if it's already reported.
2. Open a new issue with:
   - Clear title and description.
   - Steps to reproduce.
   - Logs, screenshots, or JSON samples (if applicable).

Detailed reports make fixing things much faster.

---

##  Submitting Pull Requests

When you're ready to submit code:
1. Clone the repo and create your feature branch (`git checkout -b feat/<your_feature>`).
2. Follow Go conventions (`go fmt`, `golint`).
3. Keep PRs focused on one change.
4. Add tests if possible.
5. Link your PR to an issue (e.g., `Closes #42`).

We'll review your PR and give friendly feedback.

---

##  Code Style

- Stick to idiomatic Go. Run `go fmt ./...` before pushing.
- Small, focused functions are awesome.
- Document public functions and exported types.
- Keep external dependencies minimal.
- We like to use no capital letters in our documentation

If you're working on Redis integration:
- Use the official [`go-redis`](https://github.com/redis/go-redis) client.
- Please ensure to limit changes on redis data models, many components rely on them.
---

##  Who is this for?

This project welcomes:
- Telco engineers experimenting with 5G APIs.
- Cloud-native developers building edge/vertical applications.
- Researchers needing NEF and simulator tooling.

No contribution is too small. Even opening an issue to ask a question helps.

---

##  Community Standards

We're building a **friendly, inclusive** space. Please:

- Be respectful and welcoming.
- Give constructive feedback.
- Avoid toxic or negative behavior.

---

## Need help?

- Ping us in the community chat (coming soon).
- Or email the maintainers (coming soon).

---


### Thanks for helping us build the future together.

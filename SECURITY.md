# Security policy

## Supported versions

This repository is a taskForge. Security fixes are applied on a best-effort basis to the default branch. Downstream forks should track dependencies independently.

## Reporting a vulnerability

Please **do not** open a public GitHub issue for security vulnerabilities.

Instead, contact the maintainers privately (for example via GitHub Security Advisories once enabled on the repository, or a private email listed in the repo profile).

Include:

- Description of the issue and impact
- Steps to reproduce or a proof of concept
- Affected commit / tag if known

We will acknowledge the report and work on a fix or mitigation guidance.

## Hardening reminders for forks

- Never commit `.env` or real API keys
- Rotate Clerk, Resend, New Relic, and database credentials before going public with a fork that ever contained secrets
- Lock CORS origins in production
- Prefer TLS for Postgres/Redis in shared environments
- Review Clerk middleware coverage on all authenticated routes

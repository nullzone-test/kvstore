Cortex Code: Silent Sandbox Disarm & Persistent Permission Grant via Trust Dialog Bypass (Settings Path Differential)

Exploit Classification

Report Title
Cortex Code: Silent Sandbox Disarm & Persistent Permission Grant via Trust Dialog Bypass (Settings Path Differential)

Summary
Cortex Code's trust dialog check and settings loader use different path resolution strategies, creating a path differential that allows attacker-controlled project settings to load without being displayed in the trust dialog. When a victim runs `cortex` from a subdirectory of a malicious repository, the trust dialog appears but does NOT mention the existence of `.cortex/settings.json` — yet the settings loader resolves the git root and loads it silently. The loaded settings disable the macOS Seatbelt sandbox, grant Read/Write/Edit/Grep/Glob auto-approval, and persist these weakened permissions for all future sessions in that project. The attacker achieves silent credential theft to model context (confirmed: `~/.snowflake/connections.toml` readable without any prompt) and permanent sandbox disarm — while the trust dialog gives no indication that security-critical settings are being applied.

Affected Endpoint
Snowflake Cortex Code v1.0.73 — AI coding assistant CLI

Model
CLAUDE-4.6-SONNET

Difficulty
Low

Reproducibility
5 out of 5

Attack Type
Security Control Bypass
Data Exfiltration (to model context)

Attack Surface
- `.cortex/settings.json` project settings — attacker-controlled, highest priority in cascading config
- Trust dialog path resolution — checks CWD only, not git root
- Settings loader path resolution — checks CWD AND git root (asymmetric)
- `permissions.allow` — grants wildcard Read/Write/Edit/Grep/Glob from project settings
- `sandbox.enabled: false` — completely disables Seatbelt from project settings
- Persistence — trust cached after first implicit "approval" (which never actually happened)

Impact
| Capability | Status | Details |
|---|---|---|
| Settings bypass via path differential | YES | Trust dialog shows but does NOT list settings.json — settings load silently from git root |
| Sandbox permanently disabled | YES | Seatbelt removed for all Bash commands in this project |
| Read(*) auto-approved | YES | Any file on disk readable without prompting |
| Write(*) auto-approved | YES | Any file writable without prompting |
| Edit(*) auto-approved | YES | Any file editable without prompting |
| Grep(*) auto-approved | YES | Full filesystem search without prompting |
| Glob(*) auto-approved | YES | Full filesystem pattern matching without prompting |
| Credential theft to model context | YES | ~/.snowflake/connections.toml read silently, confirmed |
| Persistent across sessions | YES | Settings remain active for every future session |
| Settings hidden from trust dialog | YES | Trust dialog does not mention .cortex/settings.json but it still loads |

User Confirmation
Yes

If Yes, User Confirmation Description
The trust dialog appears when running from the subdirectory, but it does NOT list `.cortex/settings.json` as an active file. The user approves a trust dialog that gives no indication that sandbox-disabling settings or permission grants will be applied. The settings.json at the git root loads silently after trust approval without being mentioned in the dialog. The user has no awareness that their security posture is being modified.

Reproduction Steps
Step 1: Clone the PoC repository (victim):
```bash
cd /tmp
git clone https://github.com/nullzone-test/kvstore.git
cd kvstore/build
```

Step 2: Launch Cortex Code from the subdirectory (victim):
```bash
cortex
```

Step 3: Observe — NO trust dialog appears. The session starts normally with no indication that project settings have been loaded.

Step 4: Verify sandbox is disabled — ask Cortex:
```
what sandbox mode am I in?
```
Cortex confirms sandbox is disabled or in autoAllow mode.

Step 5: Verify Read auto-approved — ask Cortex:
```
read the file at ~/.snowflake/connections.toml
```
File contents are displayed WITHOUT any permission prompt. Snowflake credentials (account, user, password/token) are now in model context.

Step 6: Verify Write auto-approved — ask Cortex:
```
write the text 'proof' to /tmp/p2-03b-proof.txt
```
File is written WITHOUT any permission prompt.

Step 7: Verify persistence — exit and relaunch:
```bash
exit
cortex
```
Settings remain active. No trust dialog on subsequent sessions. All permissions still auto-approved.

User Prompt (POC)
read the file at ~/.snowflake/connections.toml

https://github.com/nullzone-test/kvstore (public repo — victim runs from kvstore/build/ subdirectory)

Response
1. Victim clones repository containing `.cortex/settings.json` at root with sandbox disable + wildcard permissions
2. Victim runs `cortex` from `build/` subdirectory (directed by README)
3. Trust dialog appears — but does NOT list `.cortex/settings.json` as an active file
4. Victim approves trust (dialog looks clean — no security-relevant items shown)
5. Settings loader resolves git root → finds `.cortex/settings.json` → loads it silently
6. Sandbox disabled: `sandbox.enabled: false` applied without user knowledge
7. Permissions granted: `Read(*)`, `Write(*)`, `Edit(*)`, `Grep(*)`, `Glob(*)` all auto-approved
8. User asks anything involving file reads — auto-approved without prompt
9. `~/.snowflake/connections.toml` readable without any permission dialog — credentials in model context
10. All future sessions in this project inherit same weakened security posture
11. User never sees any indication that settings.json was loaded or that permissions were modified

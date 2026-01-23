# Antigravity Kit Architecture

> **Version 5.0** - Comprehensive AI Agent Capability Expansion Toolkit

---

## 📋 Overview

Antigravity Kit is a modular system consisting of:
- **16 Specialist Agents** - Role-based AI personas
- **40 Skills** - Domain-specific knowledge modules
- **11 Workflows** - Slash command procedures

---

## 🏗️ Directory Structure

```
.agent/
├── ARCHITECTURE.md          # This file
├── agents/                  # 16 Specialist Agents
├── skills/                  # 40 Skills
├── workflows/               # 11 Slash Commands
├── rules/                   # Global Rules
└── .shared/                 # Shared Resources
```

---

## 🤖 Agents (16)

Specialist AI personas for different domains.

| Agent | Focus | Skills Used |
|-------|-------|-------------|
| `orchestrator` | Multi-agent coordination | parallel-agents, behavioral-modes |
| `project-planner` | Discovery, task planning | brainstorming, plan-writing, architecture |
| `frontend-specialist` | Web UI/UX | frontend-design, react-patterns, tailwind-patterns |
| `backend-specialist` | API, business logic | api-patterns, nodejs-best-practices, database-design |
| `database-architect` | Schema, SQL | database-design, prisma-expert |
| `mobile-developer` | iOS, Android, RN | mobile-design |
| `game-developer` | Game logic, mechanics | game-development |
| `devops-engineer` | CI/CD, Docker | deployment-procedures, docker-expert |
| `security-auditor` | Security compliance | vulnerability-scanner, red-team-tactics |
| `penetration-tester` | Offensive security | red-team-tactics |
| `test-engineer` | Testing strategies | testing-patterns, tdd-workflow, webapp-testing |
| `debugger` | Root cause analysis | systematic-debugging |
| `performance-optimizer` | Speed, Web Vitals | performance-profiling |
| `seo-specialist` | Ranking, visibility | seo-fundamentals, geo-fundamentals |
| `documentation-writer` | Manuals, docs | documentation-templates |
| `explorer-agent` | Codebase analysis | - |

---

## 🧠 Skills (220+)

Domain-specific knowledge modules. Skills are loaded on-demand based on task context.

### 🤖 AI & Agents
| Skill | Description |
|-------|-------------|
| `agent-evaluation` | Testing and benchmarking LLM agents |
| `agent-manager-skill` | Local CLI agent management via tmux |
| `agent-memory-systems` | Implementation of short/long-term agent memory |
| `agent-tool-builder` | Design of robust tools for AI agents |
| `ai-agents-architect` | designing and building autonomous AI agents |
| `ai-product` | Integration patterns for AI into products |
| `autonomous-agents` | Architecture for independent goal-seeking agents |
| `crewai` | Role-based multi-agent framework expert |
| `langgraph` | Stateful, multi-actor agent applications |
| `loki-mode` | Multi-agent autonomous startup system |
| `prompt-engineer` | Expert prompt design and optimization |
| `rag-engineer` | Retrieval-Augmented Generation systems expert |
| `voice-agents` | Real-time voice interaction systems |

### 🛡️ Security & Pentesting
| Skill | Description |
|-------|-------------|
| `active-directory-attacks` | AD exploitation and security assessment |
| `api-fuzzing-bug-bounty` | API security testing and fuzzing |
| `aws-penetration-testing` | AWS cloud security assessment |
| `burp-suite-testing` | Web app testing with Burp Suite |
| `ethical-hacking-methodology` | Standard pentest lifecycle and recon |
| `linux-privilege-escalation` | Linux privesc techniques |
| `metasploit-framework` | Exploitation using Metasploit |
| `red-team-tactics` | MITRE ATT&CK based adversary simulation |
| `vulnerability-scanner` | Automated vulnerability analysis |
| `web-application-testing` | OWASP top 10 and web testing |
| `windows-privilege-escalation` | Windows privesc techniques |
| `wireshark-analysis` | Network traffic packet analysis |

### 🎨 Frontend & Design
| Skill | Description |
|-------|-------------|
| `3d-web-experience` | Three.js, R3F, WebGL experiences |
| `canvas-design` | Programmatic image generation |
| `frontend-design` | Production-grade UI/UX patterns |
| `frontend-dev-guidelines` | React/TS frontend architecture guide |
| `mobile-design` | Mobile-first UI/UX decision making |
| `react-patterns` | Modern React hooks and performance |
| `tailwind-patterns` | Tailwind CSS v4 best practices |
| `ui-ux-pro-max` | Comprehensive design systems & palettes |
| `web-artifacts-builder` | Complex HTML/React artifact generation |

### ☁️ Backend & Cloud
| Skill | Description |
|-------|-------------|
| `api-patterns` | REST, GraphQL, tRPC design |
| `aws-serverless` | AWS Lambda & Serverless architecture |
| `backend-dev-guidelines` | Node/Express/TS microservices guide |
| `database-design` | Schema optimization and normalization |
| `docker-expert` | Containerization and orchestration |
| `firebase` | Full backend-as-a-service implementation |
| `nestjs-expert` | NestJS module architecture |
| `prisma-expert` | Prisma ORM and migrations |
| `server-management` | Infrastructure and process management |
| `supabase-auth` | Authentication with Supabase |

### 📈 Marketing & Growth
| Skill | Description |
|-------|-------------|
| `app-store-optimization` | ASO for iOS and Android |
| `copywriting` | Sales and marketing copy creation |
| `email-systems` | Transactional and marketing email infra |
| `launch-strategy` | Product hunt and GTM strategies |
| `marketing-ideas` | Growth tactics and channel strategy |
| `programmatic-seo` | Scalable SEO page generation |
| `seo-fundamentals` | Technical SEO and Core Web Vitals |
| `viral-generator-builder` | Viral tool and quiz creation |

### 📦 Product & Project
| Skill | Description |
|-------|-------------|
| `architecture` | System design trade-off analysis |
| `brainstorming` | Socratic requirement discovery |
| `documentation-templates` | standard READMEs and docs |
| `plan-writing` | Structured implementation planning |
| `product-manager-toolkit` | PRDs, RICE, user research |
| `project-planner` | Task breakdown and estimation |

### 🛠️ Dev Tools & Languages
| Skill | Description |
|-------|-------------|
| `bash-linux` | Terminal mastery and scripting |
| `browser-automation` | Playwright/Puppeteer scraping |
| `clean-code` | Universal coding standards |
| `git-pushing` | Conventional commits and git workflow |
| `python-patterns` | Pythonic code and framework selection |
| `systematic-debugging` | Root cause analysis methodology |
| `tdd-workflow` | Red-Green-Refactor cycle |
| `typescript-expert` | Advanced TS types and config |
| `webapp-testing` | E2E testing with Playwright |

---

## 🔄 Workflows (11)

Slash command procedures. Invoke with `/command`.

| Command | Description |
|---------|-------------|
| `/brainstorm` | Socratic discovery |
| `/create` | Create new features |
| `/debug` | Debug issues |
| `/deploy` | Deploy application |
| `/enhance` | Improve existing code |
| `/orchestrate` | Multi-agent coordination |
| `/plan` | Task breakdown |
| `/preview` | Preview changes |
| `/status` | Check project status |
| `/test` | Run tests |
| `/ui-ux-pro-max` | Design with 50 styles |

---

## 🎯 Skill Loading Protocol

```
User Request → Skill Description Match → Load SKILL.md
                                            ↓
                                    Read references/
                                            ↓
                                    Read scripts/
```

### Skill Structure

```
skill-name/
├── SKILL.md           # (Required) Metadata & instructions
├── scripts/           # (Optional) Python/Bash scripts
├── references/        # (Optional) Templates, docs
└── assets/            # (Optional) Images, logos
```

### Enhanced Skills (with scripts/references)

| Skill | Files | Coverage |
|-------|-------|----------|
| `typescript-expert` | 5 | Utility types, tsconfig, cheatsheet |
| `ui-ux-pro-max` | 27 | 50 styles, 21 palettes, 50 fonts |
| `app-builder` | 20 | Full-stack scaffolding |

---

## 📊 Statistics

| Metric | Value |
|--------|-------|
| **Total Agents** | 16 |
| **Total Skills** | 220+ |
| **Total Workflows** | 11 |
| **Coverage** | ~90% web/mobile development |

---

## 🔗 Quick Reference

| Need | Agent | Skills |
|------|-------|--------|
| Web App | `frontend-specialist` | react-patterns, nextjs-best-practices |
| API | `backend-specialist` | api-patterns, nodejs-best-practices |
| Mobile | `mobile-developer` | mobile-design |
| Database | `database-architect` | database-design, prisma-expert |
| Security | `security-auditor` | vulnerability-scanner |
| Testing | `test-engineer` | testing-patterns, webapp-testing |
| Debug | `debugger` | systematic-debugging |
| Plan | `project-planner` | brainstorming, plan-writing |

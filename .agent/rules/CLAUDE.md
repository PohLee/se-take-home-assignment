---
trigger: always_on
platform: claude
---

# CLAUDE.md - Layer 1: Core Protocol

> **Cross-Platform AI Configuration** - Mirrored with GEMINI.md  
> **Architecture**: 3-Layer Separation of Concerns  
> **Platform**: Claude AI (Cline, Windsurf, Sonnet/Opus)

---

## 🔗 CRITICAL: READ AGENT DOCUMENTATION FIRST

**MANDATORY:** Before performing ANY implementation, read the agent system documentation:

```bash
1. Read .agent/AGENTS.md         # Layer 2: Agent system overview
2. Read .agent/ARCHITECTURE.md   # System architecture
3. Identify relevant agent       # Based on task domain
4. Read .agent/agents/{agent}.md # Specific agent instructions
5. Load required skills          # From frontmatter skills: field
```

**Rule Priority**: P0 (CLAUDE.md) > P1 (Agent .md) > P2 (SKILL.md)

---

## 📥 REQUEST CLASSIFIER (STEP 1)

**Before ANY action, classify the request:**

| Request Type | Trigger Keywords | Active Layers | Result |
|--------------|------------------|---------------|--------|
| **QUESTION** | "what is", "how does", "explain" | Layer 1 only | Text Response |
| **SURVEY/INTEL**| "analyze", "list files", "overview" | Layer 1 + Explorer | Session Intel (No File) |
| **SIMPLE CODE** | "fix", "add", "change" (single file) | Layer 1 + Layer 2 (lite) | Inline Edit |
| **COMPLEX CODE**| "build", "create", "implement", "refactor" | Layer 1 + Layer 2 (full) + Agent | **Plan file Required** |
| **DESIGN/UI** | "design", "UI", "page", "dashboard" | Layer 1 + Layer 2 + Agent | **Plan file Required** |
| **SLASH CMD** | /create, /orchestrate, /debug | Command-specific flow | Variable |

---

## LAYER 1: UNIVERSAL RULES (Always Active)

### 🌐 Language Handling

When user's prompt is NOT in English:
1. **Internally translate** for better comprehension
2. **Respond in user's language** - match their communication
3. **Code comments/variables** remain in English

### 🧹 Clean Code (Global Mandatory)

**ALL code MUST follow `@[skills/clean-code]` rules. No exceptions.**

- Concise, direct, solution-focused
- No verbose explanations
- No over-commenting
- No over-engineering
- **Self-Documentation:** Every agent is responsible for documenting their own changes in relevant `.md` files.
- **Global Testing Mandate:** Every agent is responsible for writing and running tests for their changes. Follow the "Testing Pyramid" (Unit > Integration > E2E) and the "AAA Pattern" (Arrange, Act, Assert).
- **Global Performance Mandate:** "Measure first, optimize second." Every agent must ensure their changes adhere to 2025 performance standards (Core Web Vitals for Web, query optimization for DB, bundle limits for FS).
- **Infrastructure & Safety Mandate:** Every agent is responsible for the deployability and operational safety of their changes. Follow the "5-Phase Deployment Process" (Prepare, Backup, Deploy, Verify, Confirm/Rollback). Always verify environment variables and secrets security.

### 📁 File Dependency Awareness

**Before modifying ANY file:**
1. Check `CODEBASE.md` → File Dependencies (if exists)
2. Identify dependent files
3. Update ALL affected files together

### 🗺️ System Map Read

> 🔴 **MANDATORY:** Read `.agent/AGENTS.md` and `.agent/ARCHITECTURE.md` at session start to understand the 3-layer architecture.

**Path Awareness:**
- **Layer 1 (Core)**: `.agent/rules/CLAUDE.md` (this file)
- **Layer 2 (Agents)**: `.agent/agents/*.md` (16 specialist agents)
- **Layer 3 (Skills)**: `.agent/skills/*/SKILL.md` (40+ skills)
- **Workflows**: `.agent/workflows/*.md` (11 slash commands)
- **Runtime Scripts**: `.agent/skills/<skill>/scripts/*.py`

### 🧠 Read → Understand → Apply

```
❌ WRONG: Read agent file → Start coding
✅ CORRECT: Read → Understand WHY → Apply PRINCIPLES → Code
```

**Before coding, answer:**
1. What is the GOAL of this agent/skill?
2. What PRINCIPLES must I apply?
3. How does this DIFFER from generic output?

---

## LAYER 2: AGENT ROUTING (When Writing Code)

### 📱 Project Type Routing

| Project Type | Primary Agent | Skills |
|--------------|---------------|--------|
| **MOBILE** (iOS, Android, RN, Flutter) | `mobile-developer` | mobile-design |
| **WEB** (Next.js, React web) | `frontend-specialist` | frontend-design, react-patterns |
| **BACKEND** (API, server, DB) | `backend-specialist` | api-patterns, database-design |
| **TESTING** | `test-engineer` | testing-patterns, tdd-workflow |
| **DEBUGGING** | `debugger` | systematic-debugging |
| **PLANNING** | `project-planner` | brainstorming, plan-writing |

> 🔴 **Mobile + frontend-specialist = WRONG.** Mobile = mobile-developer ONLY.

### 🛑 GLOBAL SOCRATIC GATE (MANDATORY)

**MANDATORY: Every complex request must pass through the Socratic Gate before ANY tool use or implementation.**

| Request Type | Strategy | Required Action |
|--------------|----------|-----------------|
| **New Feature / Build** | Deep Discovery | ASK minimum 3 strategic questions |
| **Code Edit / Bug Fix** | Context Check | Confirm understanding + ask impact questions |
| **Vague / Simple** | Clarification | Ask Purpose, Users, and Scope |
| **Full Orchestration** | Gatekeeper | **STOP** subagents until user confirms plan details |
| **Direct "Proceed"** | Validation | **STOP** → Even if answers are given, ask 2 "Edge Case" questions |

**Protocol:** 
1. **Never Assume:** If even 1% is unclear, ASK.
2. **Handle Spec-heavy Requests:** When user gives a list (Answers 1, 2, 3...), do NOT skip the gate. Instead, ask about **Trade-offs** or **Edge Cases** (e.g., "LocalStorage confirmed, but should we handle data clearing or versioning?") before starting.
3. **Wait:** Do NOT invoke subagents or write code until the user clears the Gate.
4. **Reference:** Full protocol in `@[skills/brainstorming]`.

### 🏁 Final Checklist Protocol

**Trigger:** When the user says "final checks", "run all tests", or similar phrases.

| Task Stage | Command | Purpose |
|------------|---------|---------|
| **Manual Audit** | `python scripts/checklist.py .` | Priority-based project audit |
| **Pre-Deploy** | `python scripts/checklist.py . --url <URL>` | Full Suite + Performance + E2E |

**Priority Execution Order:**
1. **Security** → 2. **Lint** → 3. **Schema** → 4. **Tests** → 5. **UX** → 6. **SEO** → 7. **Lighthouse/E2E**

**Rules:**
- **Completion:** A task is NOT finished until `checklist.py` returns success.
- **Reporting:** If it fails, fix the **Critical** blockers first (Security/Lint).

**Available Scripts (12 total):**
| Script | Skill | When to Use |
|--------|-------|-------------|
| `security_scan.py` | vulnerability-scanner | Always on deploy |
| `dependency_analyzer.py` | vulnerability-scanner | Weekly / Deploy |
| `lint_runner.py` | lint-and-validate | Every code change |
| `test_runner.py` | testing-patterns | After logic change |
| `schema_validator.py` | database-design | After DB change |
| `ux_audit.py` | frontend-design | After UI change |
| `accessibility_checker.py` | frontend-design | After UI change |
| `seo_checker.py` | seo-fundamentals | After page change |
| `bundle_analyzer.py` | performance-profiling | Before deploy |
| `mobile_audit.py` | mobile-design | After mobile change |
| `lighthouse_audit.py` | performance-profiling | Before deploy |
| `playwright_runner.py` | webapp-testing | Before deploy |

> 🔴 **Agents & Skills can invoke ANY script** via `python .agent/skills/<skill>/scripts/<script>.py`

### 🎭 Agent Mode Mapping

| Mode | Agent | Behavior |
|------|-------|----------|
| **plan** | `project-planner` | 4-phase methodology. NO CODE before Phase 4. |
| **ask** | - | Focus on understanding. Ask questions. |
| **edit** | `orchestrator` | Execute. Check `{task-slug}.md` first. |

**Plan Mode (4-Phase):**
1. ANALYSIS → Research, questions
2. PLANNING → `{task-slug}.md`, task breakdown
3. SOLUTIONING → Architecture, design (NO CODE!)
4. IMPLEMENTATION → Code + tests

> 🔴 **Edit mode:** If multi-file or structural change → Offer to create `{task-slug}.md`. For single-file fixes → Proceed directly.

---

## LAYER 3: SKILL LOADING (Reference)

> **Skills are loaded on-demand based on agent frontmatter.**

### Modular Skill Loading Protocol
```
Agent activated → Check frontmatter "skills:" field
    │
    └── For EACH skill:
        ├── Read SKILL.md (INDEX only)
        ├── Find relevant sections from content map
        └── Read ONLY those section files
```

- **Selective Reading:** DO NOT read ALL files in a skill folder. Read `SKILL.md` first, then only read sections matching the user's request.
- **Rule Priority:** P0 (CLAUDE.md) > P1 (Agent .md) > P2 (SKILL.md). All rules are binding.

---

## 📁 QUICK REFERENCE

### Available Master Agents (6 Core)

| Agent | Domain & Focus |
|-------|----------------|
| `orchestrator` | Multi-agent coordination and synthesis |
| `project-planner` | Discovery, Architecture, and Task Planning |
| `backend-specialist` | Backend Architect (API + Database + Server/Docker Deploy) |
| `test-engineer` | Testing & QA (Unit + Integration + E2E) |
| `debugger` | Systematic Root Cause Analysis & Bug Fixing |
| `performance-optimizer` | Performance Profiling & Optimization |

### Additional Specialist Agents (10)

| Agent | Domain |
|-------|--------|
| `security-auditor` | Master Cybersecurity (Audit + Pentest + Infra Hardening) |
| `frontend-specialist` | Frontend & Growth (UI/UX + SEO + Edge/Static Deploy) |
| `mobile-developer` | Mobile Specialist (Cross-platform + Mobile Performance)|
| `game-developer` | Specialized Game Logic & Assets & Performance |
| `database-architect` | Database Design & Optimization |
| `devops-engineer` | CI/CD & Infrastructure |
| `seo-specialist` | SEO & Rankings |
| `documentation-writer` | Technical Documentation |
| `explorer-agent` | Codebase Discovery |
| `penetration-tester` | Offensive Security Testing |

### Key Skills (Layer 3)

| Skill | Purpose |
|-------|---------|
| `clean-code` | Coding standards (GLOBAL) |
| `brainstorming` | Socratic questioning |
| `app-builder` | Full-stack orchestration |
| `frontend-design` | Web UI patterns |
| `mobile-design` | Mobile UI patterns |
| `plan-writing` | {task-slug}.md format |
| `behavioral-modes` | Mode switching |
| `testing-patterns` | Test strategies |
| `api-patterns` | API design |
| `database-design` | Schema design |

### Workflow Slash Commands

| Command | Purpose |
|---------|---------|
| `/brainstorm` | Structured brainstorming |
| `/create` | Create new application |
| `/debug` | Debug issues |
| `/deploy` | Deploy to production |
| `/enhance` | Add/update features |
| `/orchestrate` | Multi-agent coordination |
| `/plan` | Create project plan |
| `/preview` | Start/stop dev server |
| `/status` | Check project status |
| `/test` | Generate/run tests |

---

## 🔗 Cross-Platform Compatibility

This configuration is **mirrored** with GEMINI.md to ensure consistent behavior across AI platforms:

- **Claude AI**: Uses this file (`.agent/rules/CLAUDE.md`)
- **Gemini AI**: Uses `.agent/rules/GEMINI.md` (identical Layer 1 rules)
- **Universal Reference**: `.agent/AGENTS.md` (Layer 2 documentation)

**All platforms share**:
- Layer 1: Core Protocol (this file)
- Layer 2: Agent system (`.agent/AGENTS.md`)
- Layer 3: Skills (`.agent/skills/*/SKILL.md`)

---

**Version**: 3-Layer Architecture v1.0  
**Platform**: Claude AI  
**Mirror**: GEMINI.md  
**Last Updated**: 2026-01-21

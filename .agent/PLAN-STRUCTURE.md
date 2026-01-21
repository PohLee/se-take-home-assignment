# Plan Folder Structure Guide

## 📋 Overview

All project plans are organized under the `plans/` directory with a standardized structure for easy tracking and execution.

## 🗂️ Folder Structure

```
plans/
├── mcdonalds-order-system/          # Example: McDonald's order management
│   ├── PLAN.md                      # Main plan file
│   ├── tasks/                       # Task breakdown (optional)
│   │   ├── 01-setup-project.md
│   │   ├── 02-order-queue.md
│   │   ├── 03-bot-management.md
│   │   └── 04-testing.md
│   └── artifacts/                   # Supporting files (optional)
│       ├── architecture.png
│       ├── api-spec.yaml
│       └── wireframes.pdf
├── user-authentication/             # Example: Auth system
│   ├── PLAN.md
│   └── tasks/
│       ├── 01-jwt-setup.md
│       └── 02-oauth-integration.md
└── performance-optimization/        # Example: Performance work
    └── PLAN.md
```

## 📝 Naming Convention

### Task Slug Format
- **Lowercase with hyphens**: `mcdonalds-order-system`
- **Descriptive**: Clearly indicates what the plan covers
- **Unique**: Each task gets its own folder

### File Naming
- **Main plan**: Always `PLAN.md` (uppercase)
- **Task files**: Numbered prefix `01-`, `02-`, etc. for execution order
- **Artifacts**: Descriptive names with appropriate extensions

## 📄 PLAN.md Template

```markdown
# {Task Title}

**Status**: Planning | In Progress | Completed  
**Created**: 2026-01-21  
**Last Updated**: 2026-01-21

## Overview
Brief description of what this plan covers.

## Goals
- [ ] Goal 1
- [ ] Goal 2
- [ ] Goal 3

## Task Breakdown
1. [Setup](tasks/01-setup.md)
2. [Feature A](tasks/02-feature-a.md)
3. [Testing](tasks/03-testing.md)

## Dependencies
- External APIs
- Libraries/frameworks
- Other plans

## Success Criteria
- Criterion 1
- Criterion 2

## Notes
Additional context, decisions, or learnings.
```

## 🔄 Workflow Integration

### Creating a Plan
```bash
# AI checks for existing plan
Read plans/{task-slug}/PLAN.md

# If not found, create structure
mkdir -p plans/{task-slug}/tasks
mkdir -p plans/{task-slug}/artifacts
touch plans/{task-slug}/PLAN.md
```

### Executing a Plan
```bash
# Read main plan
Read plans/{task-slug}/PLAN.md

# Execute tasks in order
Read plans/{task-slug}/tasks/01-*.md
Execute task 01
Mark complete

Read plans/{task-slug}/tasks/02-*.md
Execute task 02
Mark complete
```

### Updating a Plan
```bash
# Update status
Edit plans/{task-slug}/PLAN.md

# Add new task
Create plans/{task-slug}/tasks/04-new-feature.md

# Add artifact
Save plans/{task-slug}/artifacts/updated-design.png
```

## 🎯 Best Practices

### 1. One Plan Per Feature/Task
- Don't mix unrelated work in one plan
- Create separate folders for separate initiatives

### 2. Keep Tasks Atomic
- Each task file should be independently executable
- Clear inputs and outputs
- Single responsibility

### 3. Document Decisions
- Use artifacts/ for supporting materials
- Update PLAN.md with learnings
- Track status changes

### 4. Version Control
- Commit plans/ folder to git
- Track plan evolution over time
- Reference commits in task files

## 📊 Status Tracking

### Plan Status
- **Planning**: Research and design phase
- **In Progress**: Active development
- **Blocked**: Waiting on dependencies
- **Completed**: All tasks done
- **Archived**: Historical reference

### Task Status
Use checkboxes in PLAN.md:
```markdown
## Task Breakdown
- [x] 01-setup (Completed)
- [ ] 02-feature-a (In Progress)
- [ ] 03-testing (Pending)
```

## 🔍 Finding Plans

### List All Plans
```bash
ls plans/
```

### Search Plans
```bash
grep -r "keyword" plans/*/PLAN.md
```

### Recent Plans
```bash
ls -lt plans/
```

## 🗑️ Cleanup

### Archiving Old Plans
```bash
# Move completed plans to archive
mkdir -p plans/.archive
mv plans/{old-task-slug} plans/.archive/
```

### .gitignore Recommendations
```gitignore
# Keep plans in version control
# Only ignore temporary artifacts if needed
plans/**/artifacts/*.tmp
plans/**/.DS_Store
```

## 💡 Examples

### Simple Plan (No Subtasks)
```
plans/
└── fix-login-bug/
    └── PLAN.md
```

### Complex Plan (With Breakdown)
```
plans/
└── ecommerce-platform/
    ├── PLAN.md
    ├── tasks/
    │   ├── 01-database-schema.md
    │   ├── 02-product-catalog.md
    │   ├── 03-shopping-cart.md
    │   ├── 04-checkout-flow.md
    │   └── 05-payment-integration.md
    └── artifacts/
        ├── database-erd.png
        ├── api-spec.yaml
        └── user-flows.pdf
```

## 🔗 Integration with Agents

### project-planner Agent
- Creates initial `plans/{task-slug}/PLAN.md`
- Generates task breakdown in `tasks/` folder
- Updates status as work progresses

### orchestrator Agent
- Reads plan before invoking specialists
- Ensures all agents follow the plan
- Updates plan with execution results

### Specialist Agents
- Reference plan for context
- Update task files with implementation notes
- Add artifacts as needed

---

**Remember**: Plans are living documents. Update them as you learn and execute! 📝

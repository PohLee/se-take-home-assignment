# Project Plans

This directory contains all project plans organized by task slug.

## 📂 Structure

Each plan lives in its own folder:

```
plans/
├── {task-slug}/
│   ├── PLAN.md          # Main plan file (required)
│   ├── tasks/           # Task breakdown (optional)
│   └── artifacts/       # Supporting files (optional)
└── .archive/            # Completed/archived plans
```

## 🚀 Quick Start

### Creating a New Plan

1. **Use the `/plan` command** or invoke `project-planner` agent
2. AI will create: `plans/{task-slug}/PLAN.md`
3. Review and approve the plan
4. Execute tasks in order

### Executing a Plan

1. Read `plans/{task-slug}/PLAN.md`
2. Follow task breakdown
3. Update status as you progress
4. Mark complete when done

## 📋 Plan Template

See [../.agent/PLAN-STRUCTURE.md](../.agent/PLAN-STRUCTURE.md) for detailed template and best practices.

## 🔍 Finding Plans

```bash
# List all active plans
ls plans/

# Search for specific plan
grep -r "keyword" plans/*/PLAN.md

# View recent plans
ls -lt plans/
```

## 📊 Status Tracking

Plans use status indicators:
- **Planning**: Research and design
- **In Progress**: Active development
- **Blocked**: Waiting on dependencies
- **Completed**: All tasks done
- **Archived**: Moved to `.archive/`

## 🗑️ Archiving

When a plan is complete:
```bash
mv plans/{task-slug} plans/.archive/
```

## 📚 Examples

Current plans in this directory:
- (None yet - create your first plan with `/plan`)

---

**Note**: This folder is tracked in git. Plans are living documentation of project evolution.

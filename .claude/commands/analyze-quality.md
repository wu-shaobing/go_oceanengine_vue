---
name: analyze-quality
description: Analyze code quality including complexity, security, and best practices with detailed recommendations
---

# Analyze Code Quality

Analyzes the quality of code in your project, checking for security vulnerabilities, complexity issues, test coverage gaps, and best practice violations.

## Usage

```
/analyze-quality --language python --path ./src --check security complexity tests
```

## Options

### Required
- `--language`: Programming language (python, javascript, typescript, java, go)
- `--path`: Path to code directory to analyze

### Optional
- `--check`: Which checks to perform
  - `security`: Scan for vulnerabilities (default: enabled)
  - `complexity`: Analyze code complexity (default: enabled)
  - `tests`: Assess test coverage (default: enabled)
  - `practices`: Check best practices (default: enabled)
- `--output`: Output file for report (default: quality-report.json)
- `--min-score`: Minimum acceptable score (default: 0)
- `--threshold-complexity`: Max acceptable complexity (default: 10)
- `--threshold-coverage`: Min test coverage % (default: 80)

## Examples

### Full Analysis
```
/analyze-quality --language python --path ./src
```

### Security-Focused
```
/analyze-quality --language python --path ./src --check security
```

### Check Specific Directory
```
/analyze-quality --language typescript --path ./frontend/src --output frontend-quality.json
```

### Fail on Low Score
```
/analyze-quality --language python --path ./src --min-score 75
```

## What It Checks

### Security Vulnerabilities
- SQL injection patterns
- XSS vulnerabilities
- Hardcoded secrets (API keys, passwords)
- Code injection (eval, exec usage)
- Insecure random number generation

### Code Complexity
- Cyclomatic complexity of functions
- Function length analysis
- Nested complexity levels
- Code duplication

### Test Coverage
- Files without tests
- Coverage percentage
- Low-coverage files
- Test quality indicators

### Best Practices
- Error handling patterns
- Logging vs print statements
- Naming conventions
- Code organization

## Understanding Results

### Overall Score (0-100)
Weighted average across all metrics:
- Security: 35% weight
- Complexity: 25% weight
- Test Coverage: 20% weight
- Best Practices: 20% weight

### Grade Scale
- **A (90-100)**: Excellent
- **B (80-89)**: Good
- **C (70-79)**: Acceptable
- **D (60-69)**: Poor
- **F (0-59)**: Critical Issues

### Issue Severity
- 🔴 **CRITICAL**: Immediate security risk (e.g., code injection)
- 🟠 **HIGH**: Significant risk (e.g., SQL injection)
- 🟡 **MEDIUM**: Moderate issues (e.g., low test coverage)
- 🟢 **LOW**: Minor improvements (e.g., naming)

## Example Output

```
Overall Score: 78/100 (Grade: C)

Summary:
- Files Analyzed: 45
- Critical Issues: 0
- Warnings: 8
- Suggestions: 12

Critical Findings:
- No critical issues found

Top Recommendations:
1. HIGH: Fix SQL injection in user_service.py:67
2. MEDIUM: Increase test coverage to 80%+
3. LOW: Use logging instead of print()

Detailed report saved to: quality-report.json
```

## Integration

### Save Report
```
/analyze-quality --language python --path ./src --output quality-report.json
```

### Check Quality Gate
```
/analyze-quality --language python --path ./src --min-score 80
```

This will exit with error if score is below 80.

## Tips

1. **Regular Analysis**: Run after significant code changes
2. **Quality Gates**: Use `--min-score` in CI/CD
3. **Fix Priorities**: Address CRITICAL and HIGH issues first
4. **Track Progress**: Save reports to track quality over time
5. **Team Standards**: Agree on minimum score as a team

## Related Commands

- `/code-review`: Get detailed code review with this skill
- `/measure-complexity`: Focused complexity analysis
- `/security-scan`: Security-focused analysis only
- `/check-tests`: Test coverage assessment

## Skill Used

This command uses the **Code Quality Analyzer** skill located at:
```
.claude/skills/code-quality-analyzer/
```

For more details, see the skill's README.md file.

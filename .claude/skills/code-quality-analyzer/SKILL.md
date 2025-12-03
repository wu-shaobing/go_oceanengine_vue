---
name: code-quality-analyzer
description: Analyzes code quality metrics including complexity, maintainability, test coverage, and security vulnerabilities with actionable recommendations
---

# Code Quality Analyzer Skill

Analyzes code quality metrics and provides actionable recommendations for improvement. This skill performs static code analysis, checks for security vulnerabilities, measures complexity, and evaluates test coverage.

## What This Skill Does

The Code Quality Analyzer skill performs comprehensive code analysis across multiple dimensions:

1. **Code Complexity Analysis**
   - Cyclomatic complexity calculation
   - Function length analysis
   - Class coupling metrics
   - Code duplication detection

2. **Security Vulnerability Scanning**
   - SQL injection patterns
   - XSS vulnerabilities
   - Hardcoded secrets and API keys
   - Insecure random number generation
   - Authentication bypass patterns

3. **Best Practices Validation**
   - Coding standards compliance (PEP 8, ESLint)
   - Proper error handling patterns
   - Resource cleanup verification
   - Naming convention adherence

4. **Test Coverage Assessment**
   - Unit test coverage percentage
   - Integration test quality
   - Test isolation and independence
   - Mock usage analysis

## When to Use This Skill

Use this skill when:
- Preparing code for production deployment
- Conducting code reviews
- Identifying technical debt
- Onboarding to a new codebase
- Performing security audits
- Planning refactoring efforts
- Setting quality gates in CI/CD

## Input Format

### Basic Analysis
```json
{
  "action": "analyze",
  "language": "python|javascript|typescript|java|go",
  "code_path": "/path/to/code/directory",
  "options": {
    "check_security": true,
    "check_complexity": true,
    "check_tests": true,
    "min_coverage_threshold": 80
  }
}
```

### Focused Analysis
```json
{
  "action": "analyze",
  "language": "python",
  "code_path": "/path/to/project",
  "focus": ["security", "complexity"], // Only these checks
  "options": {
    "output_format": "json"
  }
}
```

### Compare Against Baseline
```json
{
  "action": "compare",
  "language": "python",
  "current_path": "/path/to/current/code",
  "baseline_path": "/path/to/baseline/code",
  "options": {
    "show_improvements": true,
    "show_regressions": true
  }
}
```

## Output Format

### Analysis Results
```json
{
  "overall_score": 85,
  "grade": "B+",
  "summary": {
    "total_files": 120,
    "analyzed_files": 118,
    "critical_issues": 3,
    "warnings": 15,
    "suggestions": 22
  },
  "detailed_results": {
    "complexity": {
      "score": 82,
      "average_complexity": 4.2,
      "high_complexity_functions": [
        {
          "file": "services/user_service.py",
          "function": "process_user_data",
          "line": 145,
          "complexity": 15,
          "recommendation": "Break down into smaller functions"
        }
      ]
    },
    "security": {
      "score": 78,
      "critical_vulnerabilities": [
        {
          "type": "SQL Injection",
          "file": "api/endpoints/users.py",
          "line": 67,
          "severity": "HIGH",
          "description": "User input concatenated directly into SQL query",
          "recommendation": "Use parameterized queries",
          "code_snippet": "cursor.execute(f'SELECT * FROM users WHERE id = {user_id}')"
        }
      ],
      "potential_issues": [
        {
          "type": "Hardcoded Secret",
          "file": "config/database.py",
          "line": 12,
          "severity": "MEDIUM",
          "description": "Database password appears to be hardcoded",
          "recommendation": "Use environment variables"
        }
      ]
    },
    "test_coverage": {
      "score": 90,
      "coverage_percentage": 87,
      "uncovered_files": [
        "utils/email_sender.py",
        "services/notification_service.py"
      ],
      "low_coverage_files": [
        {
          "file": "api/middleware/auth.py",
          "coverage": 45,
          "threshold": 80
        }
      ]
    },
    "best_practices": {
      "score": 88,
      "issues": [
        {
          "type": "Naming Convention",
          "file": "models/user_model.py",
          "line": 23,
          "description": "Function name should use snake_case",
          "recommendation": "Rename 'getUserData' to 'get_user_data'"
        }
      ]
    }
  },
  "recommendations": [
    {
      "priority": "HIGH",
      "category": "Security",
      "description": "Fix SQL injection vulnerability in user_service.py:67",
      "estimated_effort": "2 hours"
    },
    {
      "priority": "MEDIUM",
      "category": "Complexity",
      "description": "Refactor process_user_data function to reduce complexity",
      "estimated_effort": "4 hours"
    },
    {
      "priority": "LOW",
      "category": "Testing",
      "description": "Increase test coverage for auth middleware to 80%+",
      "estimated_effort": "6 hours"
    }
  ],
  "technical_debt": {
    "score": 75,
    "estimated_debt_hours": 32,
    "top_debt_items": [
      {
        "description": "Legacy authentication system needs refactoring",
        "impact": "HIGH",
        "effort": "16 hours"
      },
      {
        "description": "Duplicate code in user management modules",
        "impact": "MEDIUM",
        "effort": "8 hours"
      }
    ]
  }
}
```

## How to Use

### Step 1: Prepare Your Code
Ensure your code is accessible and the skill has read permissions to the code directory.

### Step 2: Choose Analysis Type
Decide what aspects of code quality you want to analyze:
- Full analysis (all checks)
- Security-focused analysis
- Complexity analysis
- Test coverage assessment

### Step 3: Run the Analysis
Invoke the skill with your desired parameters:

```python
# Example usage
analyzer = CodeQualityAnalyzer()
results = analyzer.analyze(
    language="python",
    code_path="/path/to/project",
    options={
        "check_security": True,
        "check_complexity": True,
        "check_tests": True
    }
)
```

### Step 4: Review Results
Analyze the output:
- Check overall grade and score
- Review critical security issues first
- Address high-complexity functions
- Improve test coverage where needed
- Plan refactoring for technical debt

### Step 5: Create Action Plan
Prioritize fixes based on:
1. Security vulnerabilities (critical)
2. Complexity issues (medium)
3. Test coverage gaps (medium)
4. Best practices (low)

## Supported Languages

Currently supported:
- **Python** (version 3.7+)
  - Uses: pylint, flake8, bandit, radon
- **JavaScript/TypeScript**
  - Uses: ESLint, SonarJS
- **Java**
  - Uses: SpotBugs, PMD
- **Go**
  - Uses: go vet, golangci-lint
- **Go** (planned: Rust, C#)

## Configuration Options

### Analysis Depth
- `shallow`: Basic checks only (faster)
- `standard`: Balanced analysis (default)
- `deep`: Comprehensive analysis (slower but thorough)

### Thresholds
- `complexity_threshold`: Max acceptable cyclomatic complexity (default: 10)
- `coverage_threshold`: Minimum test coverage % (default: 80)
- `duplication_threshold`: Max allowed duplicate lines % (default: 5)

### Output Formats
- `json`: Machine-readable format
- `html`: Report with charts and graphs
- `markdown`: Human-readable report
- `csv`: Data for further analysis

## Example Scenarios

### Scenario 1: Pre-Release Quality Gate
```json
{
  "action": "analyze",
  "language": "python",
  "code_path": "/path/to/release/candidate",
  "options": {
    "min_score": 80,
    "max_critical_issues": 0,
    "min_coverage": 85
  }
}
```

### Scenario 2: Security Audit
```json
{
  "action": "analyze",
  "language": "python",
  "code_path": "/path/to/code",
  "focus": ["security"],
  "options": {
    "security_level": "strict",
    "include_owasp_checks": true
  }
}
```

### Scenario 3: Refactoring Assessment
```json
{
  "action": "analyze",
  "language": "python",
  "code_path": "/path/to/code",
  "focus": ["complexity", "duplication"],
  "options": {
    "generate_refactoring_plan": true
  }
}
```

## Best Practices

1. **Regular Analysis**: Run this skill regularly, not just before releases
2. **Set Quality Gates**: Use minimum scores in CI/CD pipelines
3. **Track Progress**: Compare scores over time to measure improvement
4. **Prioritize**: Fix security issues first, then complexity, then style
5. **Share Results**: Include quality reports in team discussions
6. **Automate**: Integrate into your development workflow
7. **Baseline**: Establish initial quality baseline before improving
8. **Document**: Keep track of quality decisions and trade-offs

## Limitations

- Cannot analyze compiled or obfuscated code
- Dynamic code generation may not be fully analyzed
- External dependencies are not checked for vulnerabilities
- Some framework-specific patterns may need manual review
- Third-party library code is not analyzed (only your code)

## Integration Tips

### CI/CD Integration
```yaml
# GitHub Actions example
- name: Code Quality Analysis
  run: |
    code-quality-analyzer \
      --language python \
      --path ./src \
      --output quality-report.json \
      --min-score 80
```

### Pre-commit Hook
```yaml
# .pre-commit-config.yaml
repos:
  - repo: local
    hooks:
      - id: code-quality-check
        name: Code Quality Check
        entry: code-quality-analyzer
        args: [--min-score, 85, --fail-on-score-below, 75]
        language: system
```

## Future Enhancements

Planned features:
- Real-time quality monitoring dashboard
- AI-powered refactoring suggestions
- Integration with popular IDEs (VS Code, IntelliJ)
- Custom rule creation and enforcement
- Team quality benchmarks and comparisons
- Integration with issue tracking systems (Jira, GitHub Issues)

## Getting Help

If you encounter issues:
1. Check the error message for specific guidance
2. Verify file paths and permissions
3. Ensure the specified language is supported
4. Review the skill's configuration options
5. Check your input format matches the specification

## Version History

- **v1.0.0**: Initial release
  - Basic complexity analysis
  - Security vulnerability scanning
  - Test coverage assessment
  - Python and JavaScript support

# Code Quality Analyzer Skill

A comprehensive skill for analyzing code quality across multiple dimensions including security, complexity, test coverage, and best practices.

## Quick Start

### Using the Skill

```python
from analyzer import CodeQualityAnalyzer

# Create analyzer instance
analyzer = CodeQualityAnalyzer()

# Analyze code
results = analyzer.analyze(
    language="python",
    code_path="/path/to/your/project",
    options={
        "check_security": True,
        "check_complexity": True,
        "check_tests": True,
        "check_best_practices": True
    }
)

# View results
print(f"Overall Score: {results['overall_score']}/100")
print(f"Grade: {results['grade']}")
```

### Command Line Usage

```bash
# Analyze a Python project
python analyzer.py --language python --path /path/to/project --output report.json

# Fail if score below 75
python analyzer.py --language python --path /path/to/project --fail-on-score-below 75
```

## Installation

### Prerequisites

- Python 3.7 or higher
- Required Python packages:
  - None (uses only standard library)

### Setup

1. Copy the `code-quality-analyzer` folder to your Claude app's skills directory
2. Ensure the Python analyzer is executable: `chmod +x analyzer.py`

## Skill Files

- **SKILL.md**: Main skill definition with YAML frontmatter and documentation
- **analyzer.py**: Core Python implementation
- **sample_input.json**: Example input for testing
- **expected_output.json**: Expected output format
- **README.md**: This file

## How It Works

### Analysis Phases

1. **File Discovery**: Scans directory for source files matching language
2. **Parsing**: Uses AST (Abstract Syntax Tree) for Python files
3. **Metric Calculation**: Computes complexity, security issues, etc.
4. **Scoring**: Applies weighted scoring across all metrics
5. **Recommendation Generation**: Creates actionable improvement plan

### Metrics Analyzed

#### Code Complexity
- **Cyclomatic Complexity**: Decision points in functions
- **Function Length**: Long functions are harder to maintain
- **Code Duplication**: Repeated code blocks

#### Security Vulnerabilities
- **SQL Injection**: User input concatenated into SQL
- **Code Injection**: Use of eval() or exec()
- **Hardcoded Secrets**: API keys, passwords in code
- **Insecure Random**: Using random.random() for security

#### Test Coverage
- **Line Coverage**: Percentage of code executed by tests
- **Missing Tests**: Files without corresponding tests
- **Test Quality**: Assessment of test isolation

#### Best Practices
- **Error Handling**: Proper exception handling patterns
- **Logging**: Use of logging vs print statements
- **Code Style**: Adherence to language conventions

## Scoring System

### Overall Score (0-100)
Weighted average of:
- Security: 35% (most critical)
- Complexity: 25%
- Test Coverage: 20%
- Best Practices: 20%

### Grade Scale
- **A (90-100)**: Excellent quality
- **B (80-89)**: Good quality
- **C (70-79)**: Acceptable with room for improvement
- **D (60-69)**: Poor quality, needs work
- **F (0-59)**: Critical issues, urgent fixes needed

### Issue Severity Levels

#### Critical
- **Security**: Code injection, authentication bypass
- **Impact**: Immediate security risk
- **Priority**: Fix immediately

#### High
- **Security**: SQL injection, XSS vulnerabilities
- **Complexity**: Functions with complexity > 20
- **Impact**: Significant risk or maintainability issues
- **Priority**: Fix within sprint

#### Medium
- **Security**: Insecure random, weak encryption
- **Coverage**: Test coverage < 70%
- **Impact**: Moderate risk
- **Priority**: Fix in next iteration

#### Low
- **Best Practices**: Naming, style, minor issues
- **Coverage**: Coverage 70-80%
- **Impact**: Minimal risk
- **Priority**: Address when convenient

## Example Reports

### High-Quality Code (Grade A)
```json
{
  "overall_score": 92,
  "grade": "A",
  "summary": {
    "critical_issues": 0,
    "warnings": 2,
    "suggestions": 5
  },
  "recommendations": [
    {
      "priority": "LOW",
      "category": "Best Practices",
      "description": "Consider using type hints for better code documentation"
    }
  ]
}
```

### Low-Quality Code (Grade F)
```json
{
  "overall_score": 45,
  "grade": "F",
  "summary": {
    "critical_issues": 3,
    "warnings": 12,
    "suggestions": 8
  },
  "recommendations": [
    {
      "priority": "HIGH",
      "category": "Security",
      "description": "Fix SQL injection vulnerability in user_service.py:45",
      "estimated_effort": "2 hours"
    }
  ]
}
```

## Integration Examples

### CI/CD Pipeline (GitHub Actions)
```yaml
name: Code Quality Check

on: [push, pull_request]

jobs:
  quality:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Install Python
        uses: actions/setup-python@v2
        with:
          python-version: 3.9
      - name: Analyze Code
        run: |
          python .claude/skills/code-quality-analyzer/analyzer.py \
            --language python \
            --path ./src \
            --output quality-report.json \
            --fail-on-score-below 75
      - name: Upload Report
        uses: actions/upload-artifact@v2
        with:
          name: quality-report
          path: quality-report.json
```

### Pre-commit Hook
```yaml
# .pre-commit-config.yaml
repos:
  - repo: local
    hooks:
      - id: code-quality
        name: Code Quality Check
        entry: python .claude/skills/code-quality-analyzer/analyzer.py
        args: [--language, python, --path, ./src, --min-score, 80]
        language: system
        pass_filenames: false
```

### Local Development
```bash
#!/bin/bash
# check-quality.sh

echo "Running code quality analysis..."
python .claude/skills/code-quality-analyzer/analyzer.py \
  --language python \
  --path ./src \
  --output quality-report.json \
  --min-score 80

if [ $? -eq 0 ]; then
  echo "✅ Code quality check passed"
else
  echo "❌ Code quality check failed"
  echo "See quality-report.json for details"
  exit 1
fi
```

## Customization

### Adjust Thresholds
Edit `analyzer.py` to change thresholds:

```python
# In _calculate_cyclomatic_complexity
if complexity > 15:  # Changed from 10
    # Add complexity issue
```

### Add New Checks
Extend the analyzer class:

```python
def _analyze_custom_metric(self, file_path, content):
    """Add your custom analysis here."""
    # Your custom logic
    pass
```

### Language Support
Add support for new languages:

```python
def _get_source_files(self, code_path: Path, language: str):
    extensions = {
        "rust": [".rs"],
        "php": [".php"],
        # Add new languages
    }
```

## Best Practices

1. **Run Regularly**: Don't wait for release - run after each significant change
2. **Set Quality Gates**: Enforce minimum scores in CI/CD
3. **Track Trends**: Keep historical quality metrics
4. **Fix Incrementally**: Address high-priority issues first
5. **Document Decisions**: Note why certain issues are accepted
6. **Team Alignment**: Agree on quality standards as a team
7. **Automated Checks**: Integrate into developer workflow
8. **Review Reports**: Regularly review quality trends

## Limitations

- **Python**: Full AST-based analysis
- **JavaScript/TypeScript**: Basic pattern matching (not full AST)
- **Other Languages**: Limited support, expanding gradually
- **Dynamic Code**: May not fully analyze dynamically generated code
- **External Dependencies**: Only analyzes your code, not dependencies

## Troubleshooting

### Common Issues

**Issue**: "File not found"
**Solution**: Ensure the code path exists and is accessible

**Issue**: "Syntax error"
**Solution**: Fix syntax errors in your code before analysis

**Issue**: "Permission denied"
**Solution**: Ensure read permissions on code directory

**Issue**: "No source files found"
**Solution**: Verify language matches file extensions

### Getting Help

1. Check the error message for specific guidance
2. Verify input format matches specification
3. Ensure all required fields are present
4. Check file paths and permissions
5. Review the skill documentation

## Contributing

To improve this skill:

1. Add new analysis checks
2. Support additional languages
3. Improve scoring algorithms
4. Enhance reporting format
5. Add visualization features

## Version History

- **v1.0.0**: Initial release
  - Python AST-based analysis
  - Security vulnerability scanning
  - Complexity metrics
  - Test coverage assessment
  - Basic JavaScript/TypeScript support

## License

This skill is provided as-is for use with Claude Code.

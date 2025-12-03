"""
Code Quality Analyzer Skill Implementation

This module provides the core functionality for analyzing code quality,
including complexity metrics, security vulnerabilities, and best practices.
"""

import os
import json
import ast
import re
from pathlib import Path
from typing import Dict, List, Any, Optional
from dataclasses import dataclass, asdict
import subprocess
import tempfile


@dataclass
class ComplexityIssue:
    """Represents a code complexity issue."""
    file: str
    line: int
    function: str
    complexity: int
    message: str
    recommendation: str


@dataclass
class SecurityIssue:
    """Represents a security vulnerability."""
    type: str
    file: str
    line: int
    severity: str  # LOW, MEDIUM, HIGH, CRITICAL
    description: str
    recommendation: str
    code_snippet: str


@dataclass
class TestCoverage:
    """Represents test coverage metrics."""
    file: str
    coverage_percentage: float
    covered_lines: int
    total_lines: int
    uncovered_lines: List[int]


class CodeQualityAnalyzer:
    """Main analyzer class for code quality assessment."""

    def __init__(self):
        """Initialize the analyzer."""
        self.results = {
            "overall_score": 0,
            "grade": "F",
            "summary": {},
            "complexity": {},
            "security": {},
            "test_coverage": {},
            "best_practices": {},
            "recommendations": [],
            "technical_debt": {}
        }

    def analyze(self, language: str, code_path: str, options: Dict[str, Any]) -> Dict[str, Any]:
        """
        Analyze code quality for a given directory.

        Args:
            language: Programming language (python, javascript, etc.)
            code_path: Path to the code directory
            options: Analysis options

        Returns:
            Dictionary containing analysis results
        """
        # Normalize path
        code_path = Path(code_path).resolve()

        if not code_path.exists():
            raise FileNotFoundError(f"Code path not found: {code_path}")

        # Initialize results
        self.results = {
            "language": language,
            "code_path": str(code_path),
            "overall_score": 0,
            "grade": "F",
            "summary": {
                "total_files": 0,
                "analyzed_files": 0,
                "critical_issues": 0,
                "warnings": 0,
                "suggestions": 0
            },
            "complexity": {
                "score": 0,
                "average_complexity": 0,
                "high_complexity_functions": [],
                "total_functions": 0
            },
            "security": {
                "score": 100,
                "critical_vulnerabilities": [],
                "potential_issues": [],
                "total_issues": 0
            },
            "test_coverage": {
                "score": 100,
                "coverage_percentage": 0,
                "uncovered_files": [],
                "low_coverage_files": [],
                "total_files": 0
            },
            "best_practices": {
                "score": 100,
                "issues": [],
                "total_issues": 0
            },
            "recommendations": [],
            "technical_debt": {
                "score": 100,
                "estimated_debt_hours": 0,
                "top_debt_items": []
            }
        }

        # Get all source files
        source_files = self._get_source_files(code_path, language)
        self.results["summary"]["total_files"] = len(source_files)

        if not source_files:
            return self.results

        # Analyze each file
        for file_path in source_files:
            self._analyze_file(file_path, options)

        # Calculate overall scores
        self._calculate_scores(options)

        # Generate recommendations
        self._generate_recommendations(options)

        return self.results

    def _get_source_files(self, code_path: Path, language: str) -> List[Path]:
        """Get all source files for the given language."""
        extensions = {
            "python": [".py"],
            "javascript": [".js"],
            "typescript": [".ts", ".tsx"],
            "java": [".java"],
            "go": [".go"]
        }

        lang_exts = extensions.get(language.lower(), [])
        if not lang_exts:
            return []

        files = []
        for ext in lang_exts:
            files.extend(code_path.rglob(f"*{ext}"))

        # Filter out test files, __pycache__, node_modules, etc.
        exclude_patterns = [
            "__pycache__",
            ".git",
            "node_modules",
            ".pytest_cache",
            "venv",
            ".venv",
            "env",
            "dist",
            "build",
            ".next",
            ".nuxt"
        ]

        filtered_files = []
        for file in files:
            if not any(pattern in str(file) for pattern in exclude_patterns):
                filtered_files.append(file)

        return filtered_files

    def _analyze_file(self, file_path: Path, options: Dict[str, Any]):
        """Analyze a single file."""
        self.results["summary"]["analyzed_files"] += 1

        if file_path.suffix == ".py":
            self._analyze_python_file(file_path, options)
        elif file_path.suffix in [".js", ".ts", ".tsx"]:
            self._analyze_js_file(file_path, options)

    def _analyze_python_file(self, file_path: Path, options: Dict[str, Any]):
        """Analyze a Python file."""
        try:
            with open(file_path, 'r', encoding='utf-8') as f:
                content = f.read()

            # Parse AST for complexity analysis
            tree = ast.parse(content)

            # Check complexity
            if options.get("check_complexity", True):
                self._analyze_complexity(file_path, tree)

            # Check security
            if options.get("check_security", True):
                self._analyze_security_python(file_path, content)

            # Check test coverage
            if options.get("check_tests", True):
                self._analyze_test_coverage_python(file_path, content)

            # Check best practices
            if options.get("check_best_practices", True):
                self._analyze_best_practices_python(file_path, tree)

        except SyntaxError as e:
            self._add_security_issue(
                type="Syntax Error",
                file=str(file_path),
                line=e.lineno or 0,
                severity="HIGH",
                description=f"File contains syntax error: {e.msg}",
                recommendation="Fix syntax error before analysis",
                code_snippet=""
            )
        except Exception as e:
            print(f"Error analyzing {file_path}: {e}")

    def _analyze_complexity(self, file_path: Path, tree: ast.AST):
        """Analyze code complexity using AST."""
        for node in ast.walk(tree):
            if isinstance(node, ast.FunctionDef):
                complexity = self._calculate_cyclomatic_complexity(node)
                if complexity > 10:  # Threshold
                    self._add_complexity_issue(
                        file=str(file_path),
                        line=node.lineno,
                        function=node.name,
                        complexity=complexity,
                        message=f"Function has high complexity: {complexity}",
                        recommendation="Break down into smaller functions"
                    )

    def _calculate_cyclomatic_complexity(self, node: ast.FunctionDef) -> int:
        """
        Calculate cyclomatic complexity of a function.
        Based on number of decision points.
        """
        complexity = 1  # Base complexity

        for child in ast.walk(node):
            if isinstance(child, (ast.If, ast.While, ast.For, ast.AsyncFor)):
                complexity += 1
            elif isinstance(child, ast.Try):
                complexity += len(child.handlers)
            elif isinstance(child, (ast.And, ast.Or)):
                complexity += 1
            elif isinstance(child, ast.comprehension):
                complexity += 1

        return complexity

    def _analyze_security_python(self, file_path: Path, content: str):
        """Analyze Python file for security issues."""
        lines = content.split('\n')

        # Check for SQL injection
        sql_patterns = [
            r'cursor\.execute\s*\(\s*f["\'].*\{.*\}.*["\']',  # f-string interpolation
            r'cursor\.execute\s*\(\s*["\'].*\s*\+\s*',  # String concatenation
            r'execute\s*\(\s*f["\'].*\{.*\}.*["\']',
        ]

        for i, line in enumerate(lines, 1):
            for pattern in sql_patterns:
                if re.search(pattern, line):
                    self._add_security_issue(
                        type="SQL Injection",
                        file=str(file_path),
                        line=i,
                        severity="HIGH",
                        description="Possible SQL injection vulnerability",
                        recommendation="Use parameterized queries with placeholders",
                        code_snippet=line.strip()
                    )

        # Check for hardcoded secrets
        secret_patterns = [
            (r'(?i)(password\s*=\s*["\'][^"\']+["\'])', "Hardcoded password"),
            (r'(?i)(api_key\s*=\s*["\'][^"\']+["\'])', "Hardcoded API key"),
            (r'(?i)(secret\s*=\s*["\'][^"\']+["\'])', "Hardcoded secret"),
            (r'(?i)(token\s*=\s*["\'][^"\']+["\'])', "Hardcoded token"),
        ]

        for i, line in enumerate(lines, 1):
            for pattern, desc in secret_patterns:
                if re.search(pattern, line):
                    self._add_security_issue(
                        type="Hardcoded Secret",
                        file=str(file_path),
                        line=i,
                        severity="MEDIUM",
                        description=desc,
                        recommendation="Use environment variables",
                        code_snippet=line.strip()
                    )

        # Check for eval() usage
        if re.search(r'\beval\s*\(', content):
            for i, line in enumerate(lines, 1):
                if 'eval(' in line:
                    self._add_security_issue(
                        type="Code Injection",
                        file=str(file_path),
                        line=i,
                        severity="CRITICAL",
                        description="Use of eval() can lead to code injection",
                        recommendation="Avoid eval(), use ast.literal_eval() for safe parsing",
                        code_snippet=line.strip()
                    )

        # Check for random (not secure)
        if re.search(r'\brandom\.random\(\)', content):
            for i, line in enumerate(lines, 1):
                if 'random.random()' in line:
                    self._add_security_issue(
                        type="Insecure Random",
                        file=str(file_path),
                        line=i,
                        severity="MEDIUM",
                        description="random.random() is not cryptographically secure",
                        recommendation="Use secrets.token_bytes() or secrets.SystemRandom()",
                        code_snippet=line.strip()
                    )

    def _analyze_test_coverage_python(self, file_path: Path, content: str):
        """Analyze Python file for test coverage."""
        # Check if file is a test file
        if "test" in file_path.stem or file_path.stem.startswith("test_"):
            return

        # Count lines of code
        total_lines = len([line for line in content.split('\n') if line.strip() and not line.strip().startswith('#')])

        # This is a simplified check - in reality, you'd need coverage.py
        # For now, we'll assume low coverage if no corresponding test file
        test_file = file_path.parent / f"test_{file_path.stem}.py"
        if not test_file.exists() and not any(file_path.stem in str(f) for f in file_path.parent.glob("test_*.py")):
            # Mark as uncovered
            self.results["test_coverage"]["uncovered_files"].append(str(file_path))

    def _analyze_best_practices_python(self, file_path: Path, tree: ast.AST):
        """Analyze Python file for best practices."""
        for node in ast.walk(tree):
            # Check for bare except clauses
            if isinstance(node, ast.Try):
                if not node.handlers:
                    self._add_best_practice_issue(
                        file=str(file_path),
                        line=node.lineno,
                        type="Error Handling",
                        description="Bare try without except clause",
                        recommendation="Add specific exception handling"
                    )

            # Check for print statements (should use logging)
            if isinstance(node, ast.Expr) and isinstance(node.value, ast.Call):
                if isinstance(node.value.func, ast.Name) and node.value.func.id == "print":
                    self._add_best_practice_issue(
                        file=str(file_path),
                        line=node.lineno,
                        type="Logging",
                        description="Use of print() instead of logging",
                        recommendation="Use Python's logging module"
                    )

    def _analyze_js_file(self, file_path: Path, options: Dict[str, Any]):
        """Analyze JavaScript/TypeScript file (placeholder for now)."""
        # Similar implementation for JS/TS
        # Would use tools like ESLint, TypeScript compiler API
        pass

    def _add_complexity_issue(self, file: str, line: int, function: str,
                            complexity: int, message: str, recommendation: str):
        """Add a complexity issue to results."""
        issue = ComplexityIssue(
            file=file,
            line=line,
            function=function,
            complexity=complexity,
            message=message,
            recommendation=recommendation
        )
        self.results["complexity"]["high_complexity_functions"].append(asdict(issue))
        self.results["summary"]["warnings"] += 1

    def _add_security_issue(self, type: str, file: str, line: int, severity: str,
                          description: str, recommendation: str, code_snippet: str):
        """Add a security issue to results."""
        issue = SecurityIssue(
            type=type,
            file=file,
            line=line,
            severity=severity,
            description=description,
            recommendation=recommendation,
            code_snippet=code_snippet
        )

        if severity == "CRITICAL":
            self.results["security"]["critical_vulnerabilities"].append(asdict(issue))
            self.results["summary"]["critical_issues"] += 1
        else:
            self.results["security"]["potential_issues"].append(asdict(issue))
            self.results["summary"]["warnings"] += 1

        self.results["security"]["total_issues"] += 1

        # Lower security score based on severity
        severity_scores = {"LOW": 95, "MEDIUM": 85, "HIGH": 70, "CRITICAL": 50}
        self.results["security"]["score"] = min(
            self.results["security"]["score"],
            severity_scores.get(severity, 90)
        )

    def _add_best_practice_issue(self, file: str, line: int, type: str,
                                description: str, recommendation: str):
        """Add a best practices issue."""
        issue = {
            "type": type,
            "file": file,
            "line": line,
            "description": description,
            "recommendation": recommendation
        }
        self.results["best_practices"]["issues"].append(issue)
        self.results["best_practices"]["total_issues"] += 1
        self.results["summary"]["suggestions"] += 1

        # Lower best practices score
        self.results["best_practices"]["score"] = max(
            50,
            self.results["best_practices"]["score"] - 5
        )

    def _calculate_scores(self, options: Dict[str, Any]):
        """Calculate overall scores based on all metrics."""
        # Calculate complexity score
        high_complexity_count = len(self.results["complexity"]["high_complexity_functions"])
        if high_complexity_count == 0:
            self.results["complexity"]["score"] = 100
        elif high_complexity_count < 5:
            self.results["complexity"]["score"] = 80
        elif high_complexity_count < 10:
            self.results["complexity"]["score"] = 70
        else:
            self.results["complexity"]["score"] = 60

        # Calculate test coverage score
        uncovered_files = self.results["test_coverage"]["uncovered_files"]
        total_files = self.results["summary"]["total_files"]
        if total_files > 0:
            coverage_pct = ((total_files - len(uncovered_files)) / total_files) * 100
            self.results["test_coverage"]["coverage_percentage"] = coverage_pct
            self.results["test_coverage"]["score"] = coverage_pct

        # Calculate weighted overall score
        weights = {
            "complexity": 0.25,
            "security": 0.35,
            "test_coverage": 0.20,
            "best_practices": 0.20
        }

        overall = (
            self.results["complexity"]["score"] * weights["complexity"] +
            self.results["security"]["score"] * weights["security"] +
            self.results["test_coverage"]["score"] * weights["test_coverage"] +
            self.results["best_practices"]["score"] * weights["best_practices"]
        )

        self.results["overall_score"] = int(overall)

        # Determine grade
        if overall >= 90:
            self.results["grade"] = "A"
        elif overall >= 80:
            self.results["grade"] = "B"
        elif overall >= 70:
            self.results["grade"] = "C"
        elif overall >= 60:
            self.results["grade"] = "D"
        else:
            self.results["grade"] = "F"

    def _generate_recommendations(self, options: Dict[str, Any]):
        """Generate prioritized recommendations based on issues found."""
        recommendations = []

        # Security recommendations
        for vuln in self.results["security"]["critical_vulnerabilities"]:
            recommendations.append({
                "priority": "HIGH",
                "category": "Security",
                "description": f"Fix {vuln['type'].lower()} vulnerability in {os.path.basename(vuln['file'])}:{vuln['line']}",
                "estimated_effort": "2-4 hours",
                "severity": vuln['severity']
            })

        # Complexity recommendations
        for issue in self.results["complexity"]["high_complexity_functions"]:
            recommendations.append({
                "priority": "MEDIUM",
                "category": "Complexity",
                "description": f"Refactor {issue['function']} function (complexity: {issue['complexity']})",
                "estimated_effort": "3-6 hours",
                "severity": "MEDIUM"
            })

        # Test coverage recommendations
        if self.results["test_coverage"]["coverage_percentage"] < 80:
            recommendations.append({
                "priority": "MEDIUM",
                "category": "Testing",
                "description": f"Increase test coverage from {self.results['test_coverage']['coverage_percentage']:.1f}% to 80%+",
                "estimated_effort": "Varies",
                "severity": "MEDIUM"
            })

        # Sort by priority
        priority_order = {"HIGH": 0, "MEDIUM": 1, "LOW": 2}
        recommendations.sort(key=lambda x: priority_order.get(x["priority"], 3))

        self.results["recommendations"] = recommendations[:10]  # Top 10 recommendations

        # Calculate technical debt
        high_priority_count = sum(1 for r in recommendations if r["priority"] == "HIGH")
        medium_priority_count = sum(1 for r in recommendations if r["priority"] == "MEDIUM")

        estimated_debt = high_priority_count * 4 + medium_priority_count * 2
        self.results["technical_debt"]["estimated_debt_hours"] = estimated_debt
        self.results["technical_debt"]["score"] = max(50, 100 - estimated_debt * 2)


def main():
    """Main entry point for command-line usage."""
    import argparse

    parser = argparse.ArgumentParser(description="Code Quality Analyzer")
    parser.add_argument("--language", required=True, help="Programming language")
    parser.add_argument("--path", required=True, help="Path to code directory")
    parser.add_argument("--output", default="quality-report.json", help="Output file")
    parser.add_argument("--min-score", type=int, default=0, help="Minimum acceptable score")
    parser.add_argument("--fail-on-score-below", type=int, default=0, help="Exit with error if score below this value")

    args = parser.parse_args()

    analyzer = CodeQualityAnalyzer()
    results = analyzer.analyze(
        language=args.language,
        code_path=args.path,
        options={
            "check_security": True,
            "check_complexity": True,
            "check_tests": True,
            "check_best_practices": True
        }
    )

    # Write results
    with open(args.output, 'w') as f:
        json.dump(results, f, indent=2)

    print(f"\nOverall Score: {results['overall_score']}/100 (Grade: {results['grade']})")
    print(f"Critical Issues: {results['summary']['critical_issues']}")
    print(f"Warnings: {results['summary']['warnings']}")
    print(f"Suggestions: {results['summary']['suggestions']}")
    print(f"\nDetailed report saved to: {args.output}")

    # Check thresholds
    if results['overall_score'] < args.fail_on_score_below:
        print(f"\n❌ FAILED: Score {results['overall_score']} is below threshold {args.fail_on_score_below}")
        exit(1)

    if results['overall_score'] < args.min_score:
        print(f"\n⚠️ WARNING: Score {results['overall_score']} is below recommended minimum {args.min_score}")
        exit(0)


if __name__ == "__main__":
    main()

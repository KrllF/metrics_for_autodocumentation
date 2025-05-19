api_keyy = "ТВОЙ АПИ КЛЮЧ"

from openai import OpenAI
import json
import time
from pathlib import Path
import logging
from typing import Dict, Any, List

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

class DocumentationEvaluator:
    def __init__(self, api_key: str):
        self.client = OpenAI(api_key=api_key)
        self.metrics_config = {
            "weights": {
                "completeness_score": 0.3,
                "accuracy_score": 0.25,
                "style_score": 0.15,
                "examples_score": 0.1,
                "readability_score": 0.2
            },
            "thresholds": {
                "critical_score": 5.0,
                "max_retries": 3
            }
        }

    def _read_file(self, path: str) -> str:
        try:
            return Path(path).read_text(encoding="utf-8")
        except Exception as e:
            logger.error(f"File error: {str(e)}")
            raise

    def _build_prompt(self, code: str, doc: str) -> List[Dict[str, str]]:
        return [
            {
                "role": "system",
                "content": """Ты senior Python-инженер. Проанализируй документацию и верни JSON с:
- completeness_score (0-10): покрытие всех сущностей
- accuracy_score (0-10): соответствие коду
- style_score (0-10): PEP-257, типы
- examples_score (0-10): наличие примеров
- readability_score (0-10): ясность для новичков
- errors: список строк с критичными ошибками
- warnings: список потенциальных проблем
- suggestions: рекомендации по улучшению"""
            },
            {
                "role": "user",
                "content": f"[Код]\n{code}\n\n[Документация]\n{doc}"
            }
        ]

    def _calculate_weighted_score(self, metrics: Dict[str, Any]) -> float:
        try:
            return sum(
                self.metrics_config["weights"][k] * metrics.get(k, 0)
                for k in self.metrics_config["weights"]
            )
        except Exception as e:
            logger.error(f"Score calculation error: {str(e)}")
            return 0.0

    def _validate_metrics(self, metrics: Dict) -> bool:
        required = [
            "completeness_score",
            "accuracy_score",
            "style_score",
            "examples_score",
            "readability_score",
            "errors",
            "suggestions"
        ]
        return all(key in metrics for key in required)

    def evaluate(self, code_path: str, doc_path: str) -> Dict[str, Any]:
        try:
            code = self._read_file(code_path)
            doc = self._read_file(doc_path)
        except Exception as e:
            return {"error": str(e)}

        for attempt in range(self.metrics_config["thresholds"]["max_retries"]):
            try:
                start_time = time.time()

                response = self.client.chat.completions.create(
                    model="gpt-4-turbo",
                    messages=self._build_prompt(code, doc),
                    temperature=0.1,
                    response_format={"type": "json_object"},
                    timeout=15
                )

                metrics = json.loads(response.choices[0].message.content)

                if not self._validate_metrics(metrics):
                    raise ValueError("Invalid metrics structure")

                metrics["weighted_score"] = self._calculate_weighted_score(metrics)
                metrics["response_time"] = time.time() - start_time
                metrics["errors_count"] = len(metrics.get("errors", []))
                metrics["warnings_count"] = len(metrics.get("warnings", []))

                if metrics["weighted_score"] < self.metrics_config["thresholds"]["critical_score"]:
                    logger.warning("Low documentation score detected!")

                return metrics

            except json.JSONDecodeError:
                logger.error("Invalid JSON response, retrying...")
                time.sleep(1)
            except Exception as e:
                logger.error(f"API Error: {str(e)}")
                if attempt == self.metrics_config["thresholds"]["max_retries"] - 1:
                    return {"error": str(e)}

        return {"error": "Max retries exceeded"}


if __name__ == "__main__":
    evaluator = DocumentationEvaluator(api_key=api_keyy)

    result = evaluator.evaluate(
        code_path="/content/main.py", #Путь до файла с кодом
        doc_path="/content/doc.md" #Путь до файла с докой
    )

    print(json.dumps(result, indent=2, ensure_ascii=False))

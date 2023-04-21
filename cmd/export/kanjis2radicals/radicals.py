import datasets
import json

_DESCRIPTION = """\
Contains radical images with radicals ids from WaniKani or https://api.robanohashi.org/docs/index.html
"""

_METADATA_URL = "https://huggingface.co/datasets/martingrzzler/radicals/raw/main/radicals_metadata.jsonl"
_IMAGES_URL = "https://huggingface.co/datasets/martingrzzler/radicals/resolve/main/radicals.tar.gz"


class Radicals(datasets.GeneratorBasedBuilder):
    """Radicals dataset."""

    def _info(self):
        return datasets.DatasetInfo(
            description=_DESCRIPTION,
            features=datasets.Features(
                {
                    "radical_image": datasets.Image(),
                    "meta": {
                        "id": datasets.Value("int32"),
                        "characters": datasets.Value("string"),
                        "slug": datasets.Value("string"),
                    },
                }
            ),
            supervised_keys=None,
            homepage="https://robanohashi.org/",
        )

    def _split_generators(self, dl_manager):
        metadata_path = dl_manager.download(_METADATA_URL)
        images_path = dl_manager.download(_IMAGES_URL)
        images_iter = dl_manager.iter_archive(images_path)

        return [
            datasets.SplitGenerator(
                name=datasets.Split.TRAIN,
                gen_kwargs={
                    "metadata_path": metadata_path,
                    "images_iter": images_iter,
                },
            ),
        ]

    def _generate_examples(self, metadata_path, images_iter):
        radicals = []
        with open(metadata_path, encoding="utf-8") as f:
            for line in f:
                metadata = json.loads(line)
                radicals.append(metadata)

        for idx, (image_path, image) in enumerate(images_iter):
            yield image_path, {
                "meta": radicals[idx],
                "radical_image": image.read(),
            }

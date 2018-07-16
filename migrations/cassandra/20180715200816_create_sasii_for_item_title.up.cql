CREATE CUSTOM INDEX title_idx ON todo_items(title)
USING 'org.apache.cassandra.index.sasi.SASIIndex'
WITH OPTIONS = {
    'mode': 'CONTAINS',
    'analyzer_class': 'org.apache.cassandra.index.sasi.analyzer.StandardAnalyzer',
    'tokenization_enable_stemming': 'true',
    'tokenization_locale': 'en',
    'tokenization_skip_stop_words': 'true',
    'analyzed': 'true',
    'tokenization_normalize_lowercase': 'true'
};

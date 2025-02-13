import yaml


class PolicyLoader:
    def load(self, filename):
        with open(filename, "r") as stream:
            self.data = yaml.safe_load(stream)
            return self.data

    def get_rule_by_desc(self, desc):
        for rule in self.data["rules"]:
            if rule["description"] == desc:
                return rule

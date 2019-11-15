# typed constants for Localizable.strings
Simple go project that will create a `Translations.Swift` file with constants from a given `localizable.strings` file. No typos can occur and you will know if someone deleted a translation key. 🥳

## From (Localizable.strings)
```
"general_cancel" = "Annuler";
"general_next" = "Suivant";
"general_previous" = "Précédent";
"general__loading" = "Chargement";
"close" = "Fermer";
"pullToRefresh" = "Tirer pour rafraîchir";
```
## To (Translations.swift)
```
public struct Translations {
	static let generalCancel = "general_cancel"
	static let generalNext = "general_next"
	static let generalPrevious = "general_previous"
	static let generalLoading = "general__loading"
	static let close = "close"
	static let pullToRefresh = "pullToRefresh"
}
```

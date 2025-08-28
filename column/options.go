package column

import "github.com/gofred-io/gofred/style"

type Options func(column *Column)

func ID(id string) Options {
	return func(column *Column) {
		column.SetID(id)
	}
}

func CrossAxisAlignment(crossAxisAlignment style.AlignItemsType) Options {
	return func(column *Column) {
		if column.style.Display == nil {
			column.style.Display = &style.Display{}
		}

		column.style.Display.AlignItems = crossAxisAlignment
	}
}

func MainAxisAlignment(mainAxisAlignment style.JustifyContentType) Options {
	return func(column *Column) {
		if column.style.Display == nil {
			column.style.Display = &style.Display{}
		}

		column.style.Display.JustifyContent = mainAxisAlignment
	}
}

func Gap(columnGap int) Options {
	return func(column *Column) {
		column.style.Display.RowGap = columnGap
	}
}

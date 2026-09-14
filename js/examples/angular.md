# Angular + `@openfluke/lucy`

```ts
import { registerLucyElements } from "@openfluke/lucy/elements";
registerLucyElements();
```

```html
<lucy-board [attr.board-json]="board | json"></lucy-board>
<lucy-lpd-table [attr.board-json]="board | json" max="20"></lucy-lpd-table>
```

Or assign the property from the component:

```ts
@ViewChild('board') el!: ElementRef;
ngOnChanges() {
  this.el.nativeElement.board = this.board;
}
```

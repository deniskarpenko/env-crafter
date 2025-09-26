export type InputType =
    'text'
    | 'email'
    | 'password'
    | 'number'
    | 'tel'
    | 'url'
    | 'search'
    | 'date'
    | 'time'
    | 'datetime-local'
    | 'checkbox'
    | 'radio'
    | 'file'
    | 'hidden';

export interface InputItem {
    id: number;
    value: string;
    isValid?: boolean;
    errorMessage?: string;
}

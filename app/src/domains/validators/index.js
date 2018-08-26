export const isString = (value) => typeof value === 'string';

export const isNotEmptyString = (value) => isString(value) && value.trim();

/*
This is an interface to represent the base components of a policy
*/

export interface IPolicy {
  name: string;
  caption: string;
  description: string;
  id: string;

  getName(): string;
  getCaption(): string;
  getDescription(): string;
  getId(): string;
}

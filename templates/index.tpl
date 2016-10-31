{% extends "base.tpl" %}

{% block maincontent %}

  <p>
    {% for image in images %}
      <li>
        <img src="{{image}}" alt="">
        <a href="{{image}}" target="_blank">{{image}}</a>
      </li>
    {% endfor %}
  </p>

{% endblock %}
